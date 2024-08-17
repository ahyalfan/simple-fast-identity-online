package domain

import (
	"context"
	"golang_biomtrik_login_fido/dto"
)

// ingat seperti konsep biometrik login ida harus mengirim challage untuk memeriksa private key nya cocok gak

type Challenge struct {
	Id  string `db:"id"`
	Key string `db:"key"`
	// kita buat untuk private key yg dimasukan ada waktu expirednya
	ExpiredAt  int64 `db:"expired_at"`
	ValidateAt int64 `db:"validated_at"`
}

type ChallegeRepository interface {
	Save(ctx context.Context, challenge *Challenge) error
	Update(ctx context.Context, challenge *Challenge) error
	FindById(ctx context.Context, id string) (Challenge, error)
}

type ChallengeService interface {
	Generate(ctx context.Context) (dto.ChallengeData, error)
	Validate(ctx context.Context, req dto.ChallengeValidate) (dto.UserData, error)
}
