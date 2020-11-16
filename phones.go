package service_utils

import (
	"errors"
	"regexp"
	"strings"
)

func PhoneParse(p string) (string, error) {
	re := regexp.MustCompile(`(?m)(\d+)`)
	rs := re.FindAllString(p, -1)
	p = strings.Join(rs, "")
	str := strings.Replace(p, "-", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "+", "", -1)
	first := string(str[0])
	if first != "7" && first != "8" && len(str) != 11 {
		str = "7" + str
	}
	if first == "8" {
		str = strings.Replace(str, "8", "7", 1)
	}

	if len(str) < 11 {
		return "", errors.New("Invalid phone number")
	}
	return str, nil
}
