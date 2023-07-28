package models

import (
	"finalproj/utils/token"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Username  string    `json:"username" gorm:"not null;unique"`
		Email     string    `json:"email" gorm:"not null;unique"`
		Password  string    `json:"password" gorm:"not null"`
		RoleId    uint      `json:"roleId"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (map[string]interface{}, error) {
	var err error

	user := User{}
	err = db.Model(User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}

	err = VerifyPassword(password, user.Password)
	if err != nil {
		return nil, err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	userData := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role_id":  user.RoleId,
	}

	return map[string]interface{}{
		"token": token,
		"user":  userData,
	}, nil
}

func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	var count int64
	db.Model(&User{}).Count(&count)

	if count == 0 {
		// Jika belum ada data pengguna, user pertama akan menjadi admin
		user.RoleId = 1
	} else {
		// Jika sudah ada data pengguna, role akan diset menjadi user
		user.RoleId = 2
	}

	// turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}

	user.Password = string(hashedPassword)
	// remove space in username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	var err error = db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}
