package main

import (
	"fmt"
	"os"
)

func main() {
	createSeedData()
}

func createSeedData() {
	file, err := os.OpenFile("$PWD/../zarf/seed/20230314184150_seed_users_table.up.sql", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("[ERROR] %v", err)
	}
	defer file.Close()

}
