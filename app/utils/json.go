package utils

import (
	"encoding/json"
	"fmt"
)

func ToJson(v interface{}) string {
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonResponse)
}
