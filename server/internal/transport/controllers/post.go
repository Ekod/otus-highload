package controllers

import (
	"net/http"

	"github.com/Ekod/otus-highload/internal/dto"
	"github.com/Ekod/otus-highload/internal/transport/services"
	"github.com/Ekod/otus-highload/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type PostHandlers struct {
	Service     *services.Services
	Logger      *zap.SugaredLogger
	RedisClient *redis.Client
}

func (h *PostHandlers) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var postRequest dto.PostCreateRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		h.Logger.Errorf("[ERROR] Controllers_CreatePost - Error parsing incoming JSON with %v and err %s", postRequest, err)

		err := errors.NewHandlerBadRequestError("Invalid json request")
		c.JSON(err.Status, err)

		return
	}

	response, err := h.Service.PostService.CreatePost(ctx, postRequest.Content, postRequest.UserID)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *PostHandlers) UpdatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var postRequest dto.PostUpdateRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		h.Logger.Errorf("[ERROR] Controllers_UpdatePost - Error parsing incoming JSON with %v and err %s", postRequest, err)

		err := errors.NewHandlerBadRequestError("Invalid json request")
		c.JSON(err.Status, err)

		return
	}

	err := h.Service.PostService.UpdatePost(ctx, postRequest.Content, postRequest.PostID)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)

		return
	}

	c.Status(http.StatusOK)
}

func (h *PostHandlers) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()

	var postRequest dto.PostDeleteRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		h.Logger.Errorf("[ERROR] Controllers_DeletePost - Error parsing incoming JSON with %v and err %s", postRequest, err)

		err := errors.NewHandlerBadRequestError("Invalid json request")
		c.JSON(err.Status, err)

		return
	}

	err := h.Service.PostService.DeletePost(ctx, postRequest.PostID)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)

		return
	}

	c.Status(http.StatusOK)
}

func (h *PostHandlers) GetPost(c *gin.Context) {
	ctx := c.Request.Context()

	var postRequest dto.PostGetRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		h.Logger.Errorf("[ERROR] Controllers_GetPost - Error parsing incoming JSON with %v and err %s", postRequest, err)

		err := errors.NewHandlerBadRequestError("Invalid json request")
		c.JSON(err.Status, err)

		return
	}

	response, err := h.Service.PostService.GetPost(ctx, postRequest.PostID)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)

		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *PostHandlers) FeedPost(c *gin.Context) {
	ctx := c.Request.Context()

	var postRequest dto.PostFeedRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		h.Logger.Errorf("[ERROR] Controllers_FeedPost - Error parsing incoming JSON with %v and err %s", postRequest, err)

		err := errors.NewHandlerBadRequestError("Invalid json request")
		c.JSON(err.Status, err)

		return
	}

	response, err := h.Service.PostService.FeedPost(ctx, postRequest.FriendIDs)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)

		return
	}

	c.JSON(http.StatusOK, response)
}
