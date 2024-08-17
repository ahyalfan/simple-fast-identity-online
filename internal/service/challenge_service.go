package service

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"golang_biomtrik_login_fido/domain"
	"golang_biomtrik_login_fido/dto"
	"golang_biomtrik_login_fido/internal/util"
	"time"

	"github.com/google/uuid"
)

type challengeServiceImpl struct {
	challengeRepository domain.ChallegeRepository
	user                domain.UserRepository
}

func NewChallengeService(challengeRepository domain.ChallegeRepository, user domain.UserRepository) domain.ChallengeService {
	return &challengeServiceImpl{
		challengeRepository: challengeRepository,
		user:                user,
	}
}

func (c *challengeServiceImpl) Generate(ctx context.Context) (dto.ChallengeData, error) {
	chalenge := domain.Challenge{
		Id:        uuid.NewString(),
		Key:       util.RandomString(10),
		ExpiredAt: time.Now().Add(10 * time.Minute).Unix(), //expired timenya cuma 10 menit
	}

	err := c.challengeRepository.Save(ctx, &chalenge)
	if err != nil {
		return dto.ChallengeData{}, err
	}
	return dto.ChallengeData{
		Id:  chalenge.Id,
		Key: chalenge.Key,
	}, nil
}

func (c *challengeServiceImpl) Validate(ctx context.Context, req dto.ChallengeValidate) (dto.UserData, error) {
	challenge, err := c.challengeRepository.FindById(ctx, req.Id)
	if err != nil {
		return dto.UserData{}, err
	}
	if challenge.Id == "" {
		return dto.UserData{}, errors.New("challenge not found")
	}
	if challenge.ExpiredAt < time.Now().Unix() {
		return dto.UserData{}, errors.New("invalid challenge or expired")
	}
	// ini artinya sudah divalidasi jika sudah lebih dari 0
	if challenge.ValidateAt > 0 {
		return dto.UserData{}, errors.New("challange already validate")
	}

	// kita cari usernya berdasarkan device id
	user, err := c.user.FindByDeviceId(ctx, req.DeviceId)
	if err != nil {
		return dto.UserData{}, err
	}
	fmt.Println(user)
	if user.Id == "" {
		return dto.UserData{}, errors.New("user not found")
	}

	pubKeyBase64, _ := base64.StdEncoding.DecodeString(user.PublicKey) // ini kita ambil dari user yg sudah melakukan register
	// yg punya public keynya
	signBase64, _ := base64.StdEncoding.DecodeString(req.Sign) // ini ambil dari frondent, yg menrgirimkan private keynya dari client

	// jadi disini kita validasi public key yg sudah pernah diregistrasikan oleh cliint dari biometeriknya
	// dengan private yg akan dikirimkan nya melalui challange
	// kita validasi menggunaknan algoritma ed25519
	publicKey := ed25519.PublicKey(pubKeyBase64)
	// jadi nanti signiture key akan dibuat oleh frondent berdasarkan challenge key dan private key pengguna
	// dan kita coba cocokan public ke ini bisa gak di verifikasi dari challange dan signature
	// jadi sebenarnya private key bisa membuat banyak public key, tapi yg disiman didatabase cukup satu sudah cukup
	if !ed25519.Verify(publicKey, []byte(challenge.Key), signBase64) { // chalange key nya ini kayak data yg akan ditandatangani, kayak sebagi kontrak
		// jadi dia butuh public key yg sudah diverifikasi menggunakn private key tertentu
		return dto.UserData{}, errors.New("invalid signature")
	}
	// update challenge validateAt
	challenge.ValidateAt = time.Now().Unix()
	err = c.challengeRepository.Update(ctx, &challenge)
	if err != nil {
		return dto.UserData{}, err
	}

	return dto.UserData{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

// Ed25519 adalah algoritma kriptografi yang dirancang khusus untuk tanda tangan digital.
//  Ini adalah bagian dari keluarga algoritma Edwards-Curve Digital Signature Algorithm (EdDSA) yang menggunakan kurva eliptik Curve25519.
//   Ed25519 terkenal karena keamanannya yang kuat dan performanya yang efisien.

// pub: Kunci publik (ed25519.PublicKey) yang digunakan untuk memverifikasi tanda tangan.
// msg: Data asli yang ditandatangani, dalam bentuk byte slice ([]byte).
// Ini adalah data yang Anda verifikasi untuk memastikan bahwa tanda tangan yang diberikan sesuai dengan data ini.
// sig: Tanda tangan digital dalam bentuk byte slice ([]byte).
// Ini adalah tanda tangan yang ingin Anda verifikasi, yang dihasilkan oleh kunci privat untuk data tertentu.
