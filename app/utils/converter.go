package utils

import (
	"strconv"
)

func StringToInt(req string) int {
	res, _ := strconv.Atoi(req)
	return res
}
