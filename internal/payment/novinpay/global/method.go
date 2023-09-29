package global

import "os"

func (code Code) String() string {
	return ResponseText[code]
}

func GetWsContext() WsContext {
	return WsContext{
		UserId:   os.Getenv("NOVIN_PAY_USERNAME"),
		Password: os.Getenv("NOVIN_PAY_PASSWORD"),
	}
}
