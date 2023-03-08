package utils

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/otisnado/fofo-server/models"
)

func AdminCreation() {

	var adminUser models.User
	adminUser.Name = "Admin"
	adminUser.Lastname = "System"
	adminUser.Mail = "admin@fofo-server.com"
	adminUser.Username = "root"
	adminUser.Password = "admin"
	adminUser.Group = 1
	adminUser.State = true

	if err := CheckIfMailExists(adminUser.Mail); err != nil {
		return
	}

	if err := CheckIfUsernameExists(adminUser.Username); err != nil {
		return
	}

	passwordOutput := base64.StdEncoding.EncodeToString([]byte(adminUser.Password))
	log.Println("Password encoded in base64: ", passwordOutput)

	if err := adminUser.HashPassword(adminUser.Password); err != nil {
		return
	}

	user := models.User{Name: adminUser.Name, Lastname: adminUser.Lastname, Username: adminUser.Username, Mail: adminUser.Mail, Password: adminUser.Password, Group: adminUser.Group, State: adminUser.State, CreatedAt: time.Now(), UpdatedAt: adminUser.UpdatedAt}
	models.DB.Create(&user)
}

func GroupAdminCreation() {
	var groupAdmin models.Group
	groupAdmin.Name = "Administrators"
	result := models.DB.Where("name = ?", groupAdmin.Name).Find(&groupAdmin)
	if result.RowsAffected > 0 {
		return
	}
	group := models.Group{Name: groupAdmin.Name, CreatedAt: time.Now(), UpdatedAt: groupAdmin.UpdatedAt}
	models.DB.Create(&group)

}
