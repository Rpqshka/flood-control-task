package handler

import (
	"errors"
	"strconv"
	"strings"
)

func checkParam(n, k string) (int64, int64, error) {
	if n == "" {
		n = "1"
	}

	if k == "" {
		k = "5"
	}

	if strings.Contains(n, "-") {
		return 0, 0, errors.New("invalid param n")
	}

	if strings.Contains(k, "-") {
		return 0, 0, errors.New("invalid param k")
	}

	nInt, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	kInt, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return nInt, kInt, nil
}
