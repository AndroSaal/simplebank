package usrInfo_service

import (
	"context"
	"log/slog"
)

type UserInfoService struct {
	log                   *slog.Logger
	userRepositoryHandler UserINFORepository
}

type UserINFORepository interface {
	GetUserInfo(ctx context.Context, userId int64) (bool, error)
}

type UserInfo struct {
	Id      int64
	IsAdmin bool
}

func NewUserInfoService(log *slog.Logger, userRepositoryHandler UserINFORepository) *UserInfoService {
	return &UserInfoService{
		log:                   log,
		userRepositoryHandler: userRepositoryHandler,
	}
}

func (u *UserInfoService) IsAdminById(ctx context.Context, userId int64) (bool, error) {
	isAdmin, err := u.userRepositoryHandler.GetUserInfo(ctx, userId)

	if err != nil {
		u.log.Error("Error while getting user info", "error", err)
		return false, err
	}

	return isAdmin, nil
}
