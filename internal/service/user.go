package service

import (
	"authentication/domain"
	"authentication/dto"
	"authentication/internal/util"
	"context"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo  domain.UserRepository
	cacheRepo domain.CacheRepository
}

func NewUser(userRepo domain.UserRepository, cacheRepo domain.CacheRepository) domain.UserService {
	return &userService{
		userRepo:  userRepo,
		cacheRepo: cacheRepo,
	}
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	token := util.GenerateRandomString(72)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepo.Set("users:"+token, userJson)
	return dto.AuthRes{
		Token: token,
	}, nil
}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cacheRepo.Get("users:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	var user domain.User
	_ = json.Unmarshal(data, &user)

	return dto.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Username: user.Username,
	}, nil
}
