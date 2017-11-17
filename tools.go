// Tools
package main

import (
	"encoding/json"
)

func ObjectToJson(data interface{}) string {
	rst, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
	}
	temp := string(rst)
	logger.Info(temp)
	return temp
}
