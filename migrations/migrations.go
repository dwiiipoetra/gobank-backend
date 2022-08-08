package migrations

import (
	"gobank-backend/helpers"
	"gobank-backend/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() { //insert into accounts
	db := helpers.ConnectDB()

	users := [2]interfaces.User{
		{Username: "martin", Email: "martin@gmail.com"},
		{Username: "dina", Email: "dina@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatedPassword,
		}
		db.Create(&user)

		account := &interfaces.Account{
			Type:    "daily account",
			Name:    string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 * int(i+1)),
			UserID:  user.ID,
		}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}
