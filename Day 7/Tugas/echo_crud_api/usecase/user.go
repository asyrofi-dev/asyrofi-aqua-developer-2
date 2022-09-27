package usecase

import (
	"echo_crud_api/entity"
	"echo_crud_api/repository"
)

type IUserUseCase interface {
	StoreUser(request entity.UserRequest) (entity.UserResponse, error)
	GetAllUsers() ([]entity.UserResponse, error)
	GetUserById(id int) (entity.UserResponse, error)
	UpdateUser(id int, request entity.UserRequest) (entity.UserResponse, error)
	DeleteUser(id int) (entity.MessageResponse, error)
}

type UserUseCase struct {
	UserRepository repository.IUserRepository
}

func NewUserUseCase(userRepo repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepo}
}

func (userUC UserUseCase) StoreUser(request entity.UserRequest) (entity.UserResponse, error) {
	user := entity.User{
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
	}

	result, err := userUC.UserRepository.StoreUser(user)

	if err != nil {
		return entity.UserResponse{}, err
	}

	response := entity.UserResponse(result)

	return response, nil
}

func (userUC UserUseCase) GetAllUsers() ([]entity.UserResponse, error) {
	result, err := userUC.UserRepository.GetAllUsers()

	if err != nil {
		return nil, err
	}

	var response []entity.UserResponse

	for _, user := range result {
		response = append(response, entity.UserResponse(user))
	}

	return response, nil
}

func (userUC UserUseCase) GetUserById(id int) (entity.UserResponse, error) {
	result, err := userUC.UserRepository.GetUserById(id)

	if err != nil {
		return entity.UserResponse{}, err
	}

	return entity.UserResponse(result), nil
}

func (userUC UserUseCase) UpdateUser(id int, request entity.UserRequest) (entity.UserResponse, error) {

	user, err := userUC.UserRepository.GetUserById(id)

	if err != nil {
		return entity.UserResponse{}, err
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Phone = request.Phone

	result, err := userUC.UserRepository.UpdateUser(user)

	if err != nil {
		return entity.UserResponse{}, err
	}

	response := entity.UserResponse(result)

	return response, nil
}

func (userUC UserUseCase) DeleteUser(id int) (entity.MessageResponse, error) {
	rows, deleteUser := userUC.UserRepository.DeleteUser(id)
	if err := deleteUser; err != nil {
		return entity.MessageResponse{
			Message: "Gagal Menghapus User",
		}, err
	} else if rows < 1 {
		return entity.MessageResponse{
			Message: "ID User Tidak Ditemukan",
		}, nil
	}

	return entity.MessageResponse{
		Message: "User Berhasil Dihapus",
	}, nil
}
