package conv

import "net/mail"

func ToEmailAddress(s string) (string, error) {
	addr, err := mail.ParseAddress(s)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

func IsEmailAddress(s string) bool {
	addr, _ := ToEmailAddress(s)
	return addr != ""
}
