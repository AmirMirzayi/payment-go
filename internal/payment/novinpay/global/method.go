package global

import (
	"os"
	"strings"
)

func (code Code) String() string {
	return ResponseText[code]
}

func GetWsContext() WsContext {
	return WsContext{
		UserId:   os.Getenv("NOVIN_PAY_USERNAME"),
		Password: os.Getenv("NOVIN_PAY_PASSWORD"),
	}
}

func GetPayCheck(status string) string {
	if strings.ToLower(status) == strings.ToLower(PaidStatus) {
		return PaidStatus
	}
	return CancelledStatus
}
