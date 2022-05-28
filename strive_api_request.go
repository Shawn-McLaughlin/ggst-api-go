package ggst_api

type striveApiRequest interface {
	getRoute() string
	asHexData() string
}
