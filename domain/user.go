package domain

import (
	"context"
	"golang_biomtrik_login_fido/dto"
)

type User struct {
	// sebenanrya device sama user ini dipish
	// tapi disini kita coba coba digabung saja, karena ini tutorial
	Id        string `db:"id"`
	Name      string `db:"name"`
	DeviceId  string `db:"device_id"`
	PublicKey string `db:"public_key"`
	CreatedAt int64  `db:"created_at"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByDeviceId(ctx context.Context, id string) (User, error)
}

type UserService interface {
	Register(ctx context.Context, req dto.UserRegisterRequest) error
}
