package khaancgw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Cro8ox/khaancgw/structs"
	"github.com/Cro8ox/khaancgw/utils"
)

var (
	AuthToken = utils.API{
		Url:    "/auth/token?grant_type=client_credentials",
		Method: http.MethodPost,
	}
	Statement = utils.API{
		Url:    "/statements/%s?from=%s&to=%s", // accountno, startdate, enddate
		Method: http.MethodPost,
	}
)

func (k khaanCGW) Auth() ([]byte, error) {
	postBody, _ := json.Marshal(nil)
	req, err := http.NewRequest(AuthToken.Method, k.endpoint+AuthToken.Url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(k.username, k.password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	_resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return _resp, err
}

func (k khaanCGW) HttpRequest(method, path string, body interface{}) ([]byte, error) {
	cred, err := k.Auth()
	if err != nil {
		return nil, err
	}
	var authResponse structs.AuthResponse
	if err := json.Unmarshal(cred, &authResponse); err != nil {
		return nil, err
	}
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}
	req, err := http.NewRequest(method, k.endpoint+path, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authResponse.AccessToken))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	_resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return _resp, nil
}
