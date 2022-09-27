package repository

import (
	"echo_crud_api/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	StoreUser(user entity.User) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id int) (int, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepo *UserRepository) StoreUser(user entity.User) (entity.User, error) {
	if err := userRepo.DB.Debug().Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (userRepo *UserRepository) GetAllUsers() ([]entity.User, error) {
	var result []entity.User

	if err := userRepo.DB.Debug().Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}

func (userRepo *UserRepository) GetUserById(id int) (entity.User, error) {
	user := entity.User{}

	if err := userRepo.DB.Debug().First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (userRepo *UserRepository) UpdateUser(user entity.User) (entity.User, error) {
	if err := userRepo.DB.Debug().Save(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
func (userRepo *UserRepository) DeleteUser(id int) (int, error) {
	deleteUser := userRepo.DB.Debug().Delete(&entity.User{}, id)

	if err := deleteUser.Error; err != nil {
		return 0, err
	}

	return int(deleteUser.RowsAffected), nil
}
