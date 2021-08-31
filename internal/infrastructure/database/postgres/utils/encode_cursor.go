package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

func EncodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func DecodeCursor(encodedCursor string) (time.Time, string, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return time.Time{}, "", err
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return time.Time{}, "", err
	}

	res, err := time.Parse(time.RFC3339, arrStr[0])
	if err != nil {
		return time.Time{}, "", err
	}
	uuid := arrStr[1]

	return res, uuid, nil
}
