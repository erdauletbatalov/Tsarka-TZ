package usecase

import (
	"regexp"
)

func Emails(str string) []string {
	re := regexp.MustCompile(`^Email:\s*[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return re.FindAllString(str, 1)
}

func PIN(str string) []string {
	re := regexp.MustCompile(`^PIN:\s*((0[48]|[2468][048]|[13579][26])0229[1-6]|000229[34]|\d\d((0[13578]|1[02])(0[1-9]|[12]\d|3[01])|(0[469]|11)(0[1-9]|[12]\d|30)|02(0[1-9]|1\d|2[0-8]))[1-6])\d{5}$`)
	return re.FindAllString(str, 1)
}
