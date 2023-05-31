package seeder

import (
	"log"
	"tutorial/config/db"
	"tutorial/database/migration"
	"tutorial/helper"

	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() bool {
	dateNow := helper.TimeNow()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Password"), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Seeder failed to add table users")
		return false
	}

	listUser := []map[string]interface{}{
		{
			"Email":    "superadmin@gmail.com",
			"Password": string(hashedPassword),
			"Name":     "Super Admin",
			"Role":     helper.SliceIndexStr(ListRoles, "Super Admin"),
		},
		{
			"Email":    "admin@gmail.com",
			"Password": string(hashedPassword),
			"Name":     "Admin",
			"Role":     helper.SliceIndexStr(ListRoles, "Admin"),
		},
		{
			"Email":    "user@gmail.com",
			"Password": string(hashedPassword),
			"Name":     "User",
			"Role":     helper.SliceIndexStr(ListRoles, "User"),
		},
	}

	var users []migration.User
	var userHasRoles []migration.UserHasRole

	for _, user := range listUser {
		users = append(users, migration.User{
			Email:     user["Email"].(string),
			Password:  user["Password"].(string),
			Name:      user["Name"].(string),
			CreatedAt: dateNow,
			UpdatedAt: dateNow,
		})
	}

	for key, user := range listUser {
		userHasRoles = append(userHasRoles, migration.UserHasRole{
			UserId:    key + 1,
			RoleId:    user["Role"].(int) + 1,
			CreatedAt: dateNow,
			UpdatedAt: dateNow,
		})
	}

	tx := db.DB.Begin()

	if err := tx.Create(users); err.Error != nil {
		log.Print("Seeder failed to add table users")
		tx.Rollback()
		return false
	}
	if err := tx.Create(userHasRoles); err.Error != nil {
		log.Print("Seeder failed to add table user has role")
		tx.Rollback()
		return false
	}

	tx.Commit()
	log.Print("Seeder table users successfully")
	return true
}
