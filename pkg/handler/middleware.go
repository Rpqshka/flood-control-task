package handler

import (
	"strconv"
)

func checkParam(n, k string) (int64, int64, error) {
	if n == "" {
		n = "0"
	}
	if k == "" {
		k = "0"
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
