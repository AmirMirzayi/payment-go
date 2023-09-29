package novinpay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"os"

	"golang.org/x/crypto/pkcs12"
)

var certificate []byte

func setCertificate() error {
	var err error
	certificate, err = os.ReadFile("/etc/dizintork/novinpardakht/cert.p12")
	if err != nil {
		return err
	}
	return nil
}

func signWithKey(privateKey *rsa.PrivateKey, data string) (string, error) {
	h := crypto.SHA256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(sign), err
}

func signToken(token string) (string, error) {

	err := setCertificate()
	if err != nil {
		return "", err
	}
	password := os.Getenv("NOVIN_PAY_CERTIFICATE_PASS")
	privateKey, _, err := pkcs12.Decode(certificate, password)
	if err != nil {
		return "", err
	}

	pv := privateKey.(*rsa.PrivateKey)
	sign, err := signWithKey(pv, token)
	if err != nil {
		return "", err
	}

	return sign, nil
}
