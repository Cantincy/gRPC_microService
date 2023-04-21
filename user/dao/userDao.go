package dao

import (
	"fmt"
	"pro01/grpc/user/entity"
)

func UserRegister(userId, pwd string) error {
	var user entity.User
	err := DB.Debug().Model(entity.User{}).Where("userId=?", userId).First(&user).Error
	if err == nil {
		return fmt.Errorf("userId already exist")
	}
	user.UserId, user.Pwd = userId, pwd
	err = DB.Debug().Model(entity.User{}).Create(&user).Error
	return err
}

func UserLogin(userId, pwd string) error {
	var user entity.User
	err := DB.Debug().Model(entity.User{}).Where("userId=? and pwd=?", userId, pwd).First(&user).Error
	return err
}

func GetUserInfo(userId string) (entity.User, error) {
	var user entity.User
	err := DB.Debug().Model(entity.User{}).Where("userId=?", userId).First(&user).Error
	return user, err
}

func ChangeUserInfo(userId, nickname, signature string) error {
	err := DB.Debug().Model(entity.User{}).Where("userId=?", userId).Updates(map[string]interface{}{
		"nickname":  nickname,
		"signature": signature,
	}).Error
	return err
}

func DeleteUser(userId string) error {
	err := DB.Debug().Where("userId=?", userId).Delete(&entity.User{}).Error
	return err
}
