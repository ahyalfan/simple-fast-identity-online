package dto

type ChallengeValidate struct {
	Id       string `json:"id"`
	Sign     string `json:"sign"`      // ini private keynya / atau signature key
	DeviceId string `json:"device_id"` //
}
