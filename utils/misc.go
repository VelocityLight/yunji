package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// `s` must be in RFC3339Nano format.
func ParseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, s)
}

const defaultMaxItem = 100

func ParseCSV(s string, maxItems ...int) ([]string, error) {
	if s == "" {
		return []string{}, nil
	}

	maxItem := defaultMaxItem
	if len(maxItems) != 0 {
		maxItem = maxItems[0]
	}

	res := strings.Split(s, ",")
	if len(res) > maxItem {
		return nil, errors.New("item count should <= " + strconv.Itoa(maxItem))
	}

	return res, nil
}
