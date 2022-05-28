package ggst_api

type loginResponse struct {
	Header loginResponseHeader
	Body   loginResponseBody
}

type loginResponseHeader struct {
	Token       string
	Int1        int
	RequestDate string
	Version1    string
	Version2    string
	Version3    string
	String1     string
	String2     string
}

type loginResponseBody struct {
	Int1     int
	UserInfo userInfo
}

type userInfo struct {
	StriveUserId string
	SteamName    string
	SteamId      string
	SteamIdHex   string
}
