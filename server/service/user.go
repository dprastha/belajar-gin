package service

import (
	"belajar-gin/server/repository"
	"belajar-gin/server/response"
	"database/sql"
)

type UserService struct {
	repo repository.UserRepo
}

func NewService(repo repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) FindUserByEmail(email string) *response.Response {
	user, err := u.repo.FindUserByEmail(email)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.ErrNotFound()
		}

		return response.ErrInternalServerError(err)
	}

	return response.SuccessFind(user)
}
