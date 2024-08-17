package dto

type UserRegisterRequest struct {
	Name      string `json:"name"`
	DeviceId  string `json:"deviceId"`
	PublicKey string `json:"publicKey"`
}
