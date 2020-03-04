package main

import (
	"fmt"
	"strings"
)

func rm_lead_space(old_value string) (new_value string) {
	new_value = strings.TrimLeft(old_value, " ")
	return
}

func formatBirthDate(month, day, year string) (birthday string) {
	birthday = fmt.Sprintf("%s%s/19%s", rm_lead_space(month), rm_lead_space(day), rm_lead_space(year))
	return
}

func formatDate(month, day, year string) (date string) {
	date = fmt.Sprintf("%s%s/20%s", rm_lead_space(month), rm_lead_space(day), rm_lead_space(year))
	return
}

func formatPhoneNumber(areacode, number string) (phonenumber string) {
	lastFour := number[3:]
	firstThree := number[0:3]
	phonenumber = fmt.Sprintf("%s-%s-%s", areacode, firstThree, lastFour)
	return
}
