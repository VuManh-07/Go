package services

import "github.com/VuManh-07/Go/Projects/go-ecommer-be-api/internal/repo"

type UserService struct {
	// Add fields if necessary
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserName() (string, error) {

	name := us.userRepo.GetUserByID()

	return name, nil
}
