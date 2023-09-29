package novinpay

import (
	"context"
	"testing"
)

func TestLogin(t *testing.T) {
	sessionId, err := login()
	if err == context.DeadlineExceeded {
		t.Errorf("server didn't respond in expected time")
	} else if err != nil {
		t.Error(err)
	} else {
		t.Log("logn is working properly")
		t.Log("session id is: ", sessionId)
	}
}
