package model

import "os"

type (
	LoginRequest struct {
		UserName string `json:"UserName"`
		Password string `json:"Password"`
	}

	LoginResponse struct {
		Result    string `json:"Result"`
		SessionId string `json:"SessionId"`
	}
)

func (res LoginResponse) GetResult() string {
	return res.Result
}

func GetLoginRequestBody() LoginRequest {
	return LoginRequest{
		UserName: os.Getenv("NOVIN_PAY_USERNAME"),
		Password: os.Getenv("NOVIN_PAY_PASSWORD"),
	}
}
