package ggst_api

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

func getJsonFromMsgPack(rawData []byte) (string, error) {
	var i interface{}
	unmarshalError := msgpack.Unmarshal(rawData, &i)
	if unmarshalError != nil {
		return "", unmarshalError
	}
	rawString := fmt.Sprint(i)
	return rawString, nil
}
