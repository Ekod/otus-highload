package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/Ekod/otus-highload/datasources/mysql"
	"github.com/Ekod/otus-highload/datasources/redis"
	"github.com/Ekod/otus-highload/internal/repositories"
	"github.com/Ekod/otus-highload/internal/services/friend_service"
	"github.com/Ekod/otus-highload/internal/services/post_service"
	"github.com/Ekod/otus-highload/internal/services/user_service"
	"github.com/Ekod/otus-highload/internal/transport/services"
	"github.com/Ekod/otus-highload/utils/logger"
	"github.com/ardanlabs/conf/v3"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

func main() {
	// Construct the application logger.
	log, err := logger.New("HIGHLOAD-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {

	// =========================================================================
	// GOMAXPROCS

	// Want to see what maxprocs reports.
	opt := maxprocs.Logger(log.Infof)

	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(opt); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// =========================================================================
	// Configuration

	cfg := struct {
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:5000"`
			DebugHost       string        `conf:"default:0.0.0.0:4000"`
		}
		DB struct {
			User         string `conf:"default:root"`
			Scheme       string `conf:"default:social"`
			Password     string `conf:"default:password,mask"`
			Host         string `conf:"default:localhost"`
			Port         string `conf:"default:3306"`
			Name         string `conf:"default:social"`
			MaxIdleConns int    `conf:"default:0"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
		SlaveDB struct {
			User         string `conf:"default:root"`
			Scheme       string `conf:"default:social"`
			Password     string `conf:"default:password,mask"`
			Host         string `conf:"default:localhost"`
			Port         string `conf:"default:3306"`
			Name         string `conf:"default:social"`
			MaxIdleConns int    `conf:"default:0"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
		Redis struct {
			Host     string `conf:"default:localhost"`
			Port     string `conf:"default:6379"`
			Password string `conf:"default:password,mask"`
			DB       int    `conf:"default:0"`
		}
	}{}

	const prefix = "HIGHLOAD"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// =========================================================================
	// Database Support

	// Create connectivity to the database.
	log.Infow("startup", "status", "initializing master database support", "host", cfg.DB.Host)

	db, err := mysql.Open(mysql.Config{
		User:         cfg.DB.User,
		Scheme:       cfg.DB.Scheme,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Port:         cfg.DB.Port,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Infow("shutdown", "status", "stopping master database support", "host", cfg.DB.Host)
		db.Close()
	}()

	log.Infow("startup", "status", "initializing slave database support", "host", cfg.SlaveDB.Host)

	slaveDB, err := mysql.Open(mysql.Config{
		User:         cfg.SlaveDB.User,
		Scheme:       cfg.SlaveDB.Scheme,
		Password:     cfg.SlaveDB.Password,
		Host:         cfg.SlaveDB.Host,
		Port:         cfg.SlaveDB.Port,
		Name:         cfg.SlaveDB.Name,
		MaxIdleConns: cfg.SlaveDB.MaxIdleConns,
		MaxOpenConns: cfg.SlaveDB.MaxOpenConns,
		DisableTLS:   cfg.SlaveDB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to slavedb: %w", err)
	}
	defer func() {
		log.Infow("shutdown", "status", "stopping slave database support", "host", cfg.SlaveDB.Host)
		db.Close()
	}()

	log.Infow("startup", "status", "initializing redis support", "host", cfg.Redis.Host)

	redisClient := redis.New(redis.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	defer func() {
		log.Infow("shutdown", "status", "stopping redis support", "host", cfg.Redis.Host)
		db.Close()
	}()

	// =========================================================================
	// App Starting

	log.Infow("starting service")
	defer log.Infow("shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Infow("startup", "config", out)

	// =========================================================================
	// Start API Service

	log.Infow("startup", "status", "initializing V1 API support")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	userRepo := repositories.NewUserRepository(db)
	userRepoSlave := repositories.NewUserRepository(slaveDB)
	friendRepo := repositories.NewFriendRepository(db)
	postRepo := repositories.NewPostRepository(db)

	userService := user_service.New(userRepo, userRepoSlave)
	friendService := friend_service.New(friendRepo)
	postService := post_service.New(postRepo)

	serviceLayer := services.New(userService, friendService, postService)

	// Construct the mux for the API calls.
	apiMux := APIMux(APIMuxConfig{
		Shutdown:    shutdown,
		Log:         log,
		Services:    serviceLayer,
		RedisClient: redisClient,
	})

	// Construct a server to service the requests against the mux.
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      apiMux,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- api.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdownChan:
		log.Infow("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
