package helper

import (
	"strconv"
	"strings"
	"time"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}

func IsNumberValue(stringNumber string) bool {
	_, err := strconv.Atoi(stringNumber)
	return err == nil
}

func IsDateValue(stringDate string) bool {
	_, err := time.Parse("01/02/2006", stringDate)
	return err == nil
}
