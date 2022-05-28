package ggst_api

import (
	"bytes"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

type loginRequest struct {
	Header loginRequestHeader
	Body   loginRequestBody
}

type loginRequestHeader struct {
	String1  string
	String2  string
	Int1     int
	Version  string
	Platform int
}

type loginRequestBody struct {
	Int1      int
	UserId    string
	UserIdHex string
	Int2      int
	String1   string
}

func createLoginRequest(steamId uint64) loginRequest {
	return loginRequest{
		Header: loginRequestHeader{
			String1:  "",
			String2:  "",
			Int1:     6,
			Version:  apiVersion,
			Platform: int(pcPlatform),
		},
		Body: loginRequestBody{
			Int1:      1,
			UserId:    fmt.Sprintf("%d", steamId),
			UserIdHex: fmt.Sprintf("%x", steamId),
			Int2:      256,
			String1:   "",
		},
	}
}

func (l *loginRequest) asHexData() string {

	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseArrayEncodedStructs(true)

	err := enc.Encode(l)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", buf.Bytes())
}

func (l *loginRequest) getRoute() string {
	return "/api/user/login"
}
