package usrInfo_service

import "log/slog"

type userInfoService struct {
	log                   *slog.Logger
	userRepositoryHandler UserINFORepository
}

type UserINFORepository interface {
	GetUserInfo(id int) (bool, error)
}

type UserInfo struct {
	Id      int
	IsAdmin bool
}

func NewUserInfoService(log *slog.Logger, userRepositoryHandler UserINFORepository) *userInfoService {
	return &userInfoService{
		log:                   log,
		userRepositoryHandler: userRepositoryHandler,
	}
}

func (u *userInfoService) GetUserInfo(id int) (UserInfo, error) {
	isAdmin, err := u.userRepositoryHandler.GetUserInfo(id)

	if err != nil {
		u.log.Error("Error while getting user info", "error", err)
		return UserInfo{
			Id:      -1,
			IsAdmin: false,
		}, err
	}

	return UserInfo{
		Id:      id,
		IsAdmin: isAdmin,
	}, nil
}
