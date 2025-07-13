package service

import "regexp"

const (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	loginRegex = `^[a-zA-Z0-9]{3,20}$`
	phoneRegex = `^\+[0-9]{11}$`
)

func isValid(pattern, input string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(input)
}

func isEmail(identifier string) bool {
	return isValid(emailRegex, identifier)
}

func isPhone(identifier string) bool {
	return isValid(phoneRegex, identifier)
}

func isLogin(identifier string) bool {
	return isValid(loginRegex, identifier)
}
