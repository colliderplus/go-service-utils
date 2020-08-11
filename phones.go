package service_utils

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

func PhoneParse(p string) (error, string) {
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
		return errors.New("Invalid phone number"), ""
	}
	return nil, str
}

func NameParse(name string) (error, []string) {
	re := regexp.MustCompile(`(?m)([a-zA-Zа-яА-ЯёË]+)+`)
	rs := re.FindAllString(name, -1)
	if len(rs) < 1 {
		return errors.New("InvalidName"), []string{}
	}
	return nil, rs
}

func PermalincValidate(permalink string) error {
	permalink = strings.ToLower(permalink)
	re := regexp.MustCompile(`(?m)([a-z0-9.]+)+`)
	rs := re.FindAllString(permalink, -1)
	str := strings.Join(rs, "")
	if len(str) == len(permalink) {
		return errors.New("InvalidCharactersInPermalink")
	}
	return nil
}

func ParseFeedDate(date string) (*time.Time, error) {
	param, err := time.Parse(time.RFC3339Nano, date) //"2006-01-02T15:04:05.999999-07:00", date)
	if err != nil {
		return nil, err
	}
	return &param, nil
}
