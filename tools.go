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
func ObjectToObejct(ctype interface{}) interface{} {
	var temp interface{}
	rst, err := json.Marshal(ctype)
	if err != nil {
		logger.Error(err)
		return nil
	}
	err = json.Unmarshal(rst, &temp)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return temp
}
