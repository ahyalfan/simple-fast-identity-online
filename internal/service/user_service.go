package service

import (
	"context"
	"encoding/base64"
	"errors"
	"golang_biomtrik_login_fido/domain"
	"golang_biomtrik_login_fido/dto"
	"time"

	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

// disini kita buat implementasi method yang dibutuhkan

func (u *userServiceImpl) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	exist, err := u.userRepository.FindByDeviceId(ctx, req.DeviceId)
	if err != nil {
		return err
	}
	if exist.Id != "" {
		return errors.New("device already registered")
	}

	// disini kita bikin publickeyny di menjadi base64 encoded
	publicKey, _ := base64.StdEncoding.DecodeString(req.PublicKey)
	if len(publicKey) != 32 {
		return errors.New("invalid public key")
	}

	newUser := &domain.User{
		Id:        uuid.NewString(), // bikin id nya dengan uuid
		DeviceId:  req.DeviceId,
		PublicKey: req.PublicKey,
		Name:      req.Name,
		CreatedAt: time.Now().Unix(), // kita bikin created atnya menggunkan time saat ini, yg maan unix ini akan menghasilkan int64
	}

	err = u.userRepository.Save(ctx, newUser)
	return err
}
