package ggst_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

type Client struct {
	steamId      uint64
	striveUserId string
	token        string
}

func NewClient(steamId uint64) Client {
	return Client{
		steamId: steamId,
	}
}

func (c *Client) Login() error {
	loginRequest := createLoginRequest(c.steamId)

	responseBytes, responseError := c.sendRequest(&loginRequest)
	if responseError != nil {
		return responseError
	}

	var loginResponse loginResponse
	unmarshalError := msgpack.Unmarshal(responseBytes, &loginResponse)
	if unmarshalError != nil {
		return unmarshalError
	}

	c.token = loginResponse.Header.Token
	c.striveUserId = loginResponse.Body.UserInfo.StriveUserId

	return nil
}

func (c *Client) sendRequest(striveRequest striveApiRequest) ([]byte, error) {
	requestBody := strings.NewReader("data=" + striveRequest.asHexData())
	requestUrl := "https://ggst-game.guiltygear.com" + striveRequest.getRoute()
	request, requestError := http.NewRequest(http.MethodPost, requestUrl, requestBody)
	if requestError != nil {
		return nil, requestError
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Cashe-Control", "no-cache")
	request.Header.Add("User-Agent", "Steam")

	client := http.Client{}
	response, responseError := client.Do(request)
	if responseError != nil {
		return nil, responseError
	}
	defer response.Request.Body.Close()

	responseBytes, readResponseErr := ioutil.ReadAll(response.Body)
	if readResponseErr != nil {
		return nil, readResponseErr
	}

	return responseBytes, nil
}

//TODO This is currently just for debugging purposes. Remove later
func (c *Client) PrintDetails() {
	fmt.Printf("Steam Id: %d\n", c.steamId)
	fmt.Printf("Strive User Id: %s\n", c.striveUserId)
	fmt.Printf("Access Token: %s\n", c.token)
}
