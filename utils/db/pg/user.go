package pg

import (
	"errors"
	"user/utils/db/model"
)

type UserService struct {

}

// 用户注册
func (this UserService) Register (Email ,Password string)(*model.User, error){
	user := &model.User{
		Email:Email,
		Password:Password,
	}
	db := PgCnn.Create(&user)
	if err := db.Error; err != nil{
		return user,err
	}else{
		return user,nil
	}
}

// 查询用户是否存在
func (this UserService) FindOneUser (Email, Password string) (*model.User, error){
	var user model.User
	err := PgCnn.Where("email = ? AND password = ?", Email, Password).Find(&user).Error
	if err != nil{
		return &user, err
	}else{
		if user.ID > 0{
			return &user, nil
		}else{
			return &user, nil
		}
	}
}

func (this UserService) FindUserById (Id int) (*model.User, error) {
	var user model.User
	err := PgCnn.Where("id = ?", Id).Find(&user).Error
	if err != nil{
		return &user, err
	}else{
		if user.ID >0{
			return &user, nil
		}else{
			err = errors.New("user not find")
			return &user, err
		}
	}
}

func (this UserService) UserUpdateById(Avatar, Nickname string, Id int) error{
	var user model.User
	err := PgCnn.Where("id = ?", Id).Find(&user).Error
	if err != nil{
		return err
	}else{
		if user.ID > 0 {
			user.Nickname = Nickname
			user.Avatar = Avatar
			PgCnn.Save(&user)
			return nil
		}else{
			err = errors.New("user not find")
			return err
		}
	}
}

func (this UserService) ActivateUserById(Id int) error{
	var user model.User
	err := PgCnn.Where("id = ?", Id).Find(&user).Error
	if err != nil{
		return err
	}else{
		if user.ID > 0 {
			user.IsActivate = true
			PgCnn.Save(&user)
			return nil
		}else{
			err = errors.New("user not find")
			return err
		}
	}
}

// find user by email
func (this UserService) FindUserByEmail(Email string) (*model.User, error){
	var user model.User
	err := PgCnn.Where("email = ?", Email).Find(&user).Error
	if err != nil{
		return &user, err
	}else{
		if user.ID > 0 && user.IsActivate == false{
			return &user, nil
		}else{
			err = errors.New("user not found or email is activate")
			return &user, err
		}
	}
}