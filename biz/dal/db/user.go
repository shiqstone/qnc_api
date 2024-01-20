package db

import (
	"qnc/pkg/constants"
	"qnc/pkg/errno"
)

type User struct {
	ID         int64   `json:"id"`
	UserName   string  `json:"user_name"`
	Password   string  `json:"password"`
	AvatarUrl  string  `json:"avatar_url"`
	Coin       float64 `json:"coin"`
	Email      string  `json:"email"`
	Salt       string  `json:"salt"`
	CreateTime int64   `json:"create_time"`
	UpdateTime int64   `json:"update_time"`
}

func (User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

// QueryUser query User by user_name
func QueryUser(userName string) (*User, error) {
	var user User
	if err := DB.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// QueryUserById get user in the database by user id
func QueryUserById(uid int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		return nil, err
	}
	if user == (User{}) {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return &user, nil
}

// VerifyUser verify username and password in the db
func VerifyUser(userName, password string) (int64, error) {
	var user User
	if err := DB.Where("user_name = ? AND password = ?", userName, password).Find(&user).Error; err != nil {
		return 0, err
	}
	if user.ID == 0 {
		err := errno.PasswordIsNotVerified
		return user.ID, err
	}
	return user.ID, nil
}

// CheckUserExistById find if user exists
func CheckUserExistById(uid int64) (bool, error) {
	var user User
	if err := DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		return false, err
	}
	if user == (User{}) {
		return false, nil
	}
	return true, nil
}
