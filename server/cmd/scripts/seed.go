package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Ekod/otus-highload/datasources/mysql"
	"github.com/Ekod/otus-highload/utils/security"
)

var (
	firstNameSlc = [10]string{"Павел", "Анна", "Геннадий", "Василиса", "Сергей", "Валентина", "Фёдор", "Мария", "Глеб", "Аглая"}
	lastNameSlc  = [10]string{"Муфлоний", "Каримовна", "Остапенков", "Премудрая", "Вяземский", "Агаповна", "Эстапьев", "Понская", "Жиглов", "Валенкова"}
	ageSlc       = [10]int{24, 54, 35, 12, 87, 78, 6, 51, 47, 99}
	interestsSlc = [10]string{"Горные лыжи", "Книги", "Телевизор", "Настольные игры", "Кровать", "Водка", "Планшет", "Машины", "Порно", "Религия"}
	citySlc      = [10]string{"Смоленск", "Уфа", "Нижний-Новгород", "Казань", "Москва", "Владивосток", "Верхние Васюки", "Урюпинск", "Новосибирск", "Красножопинск"}
	genderSlc    = [10]string{"male", "female", "male", "female", "male", "female", "male", "female", "male", "female"}
)

func main() {
	var cycleRange = flag.Int("n", 1, "flag sets amount of records that need to be generated for seed migration")
	flag.Parse()

	err := createSeedData(*cycleRange)
	if err != nil {
		log.Fatalln(err)
	}
}

func createSeedData(cycleRange int) error {
	db, err := mysql.Open(mysql.Config{
		User:         "root",
		Scheme:       "social",
		Password:     "password",
		Host:         "localhost",
		Port:         "6446",
		MaxIdleConns: 0,
		MaxOpenConns: 0,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	password, _ := security.HashPassword("password")

	for i := 1; i <= cycleRange; i++ {
		randIdx := rand.Intn(10)

		firstName := firstNameSlc[randIdx]
		lastName := lastNameSlc[randIdx]
		age := ageSlc[randIdx]
		gender := genderSlc[randIdx]
		interests := interestsSlc[randIdx]
		city := citySlc[randIdx]

		randNumForEmail := rand.Intn(1_000_000_000_000_000)
		email := fmt.Sprintf("test@%d.ru", randNumForEmail)

		record := "INSERT INTO users(first_name,last_name, age, gender, interests, city, email, password, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?, NOW(), NOW())"

		_, err = db.Exec(record, firstName, lastName, age, gender, interests, city, email, password)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Record number %d - done\n", i)
	}

	return nil
}
