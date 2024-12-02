package usrInfo_service

import (
	"context"
	"log/slog"

	"github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	auth_repository "github.com/AndtoSaal/simplebank/services/auth/src/repository"
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

func NewUserInfoService(
	log *slog.Logger,
	serviceConfig config.ServiceConfig,
) *UserInfoService {
	return &UserInfoService{
		log:                   log,
		userRepositoryHandler: auth_repository.NewAuthPostgresRepo(&serviceConfig.DB, log),
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
