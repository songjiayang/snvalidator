package snvalidator

import (
	"strconv"
)

// detail to check https://zh.wikipedia.org/wiki/中华人民共和国公民身份号码

var (
	checkValues = []int{
		7,
		9,
		10,
		5,
		8,
		4,
		2,
		1,
		6,
		3,
		7,
		9,
		10,
		5,
		8,
		4,
		2,
	}
)

func IsValidSN(sn string) (ok bool, err error) {
	if len(sn) != 18 {
		return
	}

	var (
		checkSum  int = 0
		code      int
		checkCode string
	)

	for i := 0; i < 17; i++ {
		code, err = strconv.Atoi(string(sn[i]))
		if err != nil {
			return
		}
		checkSum += code * checkValues[i]
	}

	cCode := (12 - (checkSum % 11)) % 11
	if cCode == 10 {
		checkCode = "X"
	} else {
		checkCode = strconv.Itoa(cCode)
	}

	ok = checkCode == string(sn[17])

	return ok, nil
}
