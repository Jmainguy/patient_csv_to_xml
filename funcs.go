package main

import (
	"fmt"
	"strings"
)

func rmLeadSpace(oldValue string) (newValue string) {
	newValue = strings.TrimLeft(oldValue, " ")
	return
}

func formatBirthDate(month, day, year string) (birthday string) {
	birthday = fmt.Sprintf("%s%s/19%s", rmLeadSpace(month), rmLeadSpace(day), rmLeadSpace(year))
	return
}

func formatDate(month, day, year string) (date string) {
	date = fmt.Sprintf("%s%s/20%s", rmLeadSpace(month), rmLeadSpace(day), rmLeadSpace(year))
	return
}

func formatPhoneNumber(areacode, number string) (phonenumber string) {
	lastFour := number[3:]
	firstThree := number[0:3]
	phonenumber = fmt.Sprintf("%s-%s-%s", areacode, firstThree, lastFour)
	return
}
