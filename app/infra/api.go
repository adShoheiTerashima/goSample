package infra

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIClient 外部APIのリクエスト設定用
// おそらくリクエストパラメータの設定をどうするか色々出てくると思うので、
// いったん仮置で
type APIClient struct {
	url          string
	data         interface{}
	responseType interface{}
}

// Post post
func (a *APIClient) Post() error {

	postData, parseErr := json.Marshal(a.data)
	if parseErr != nil {
		return parseErr
	}

	resp, postErr := http.Post(a.url, "application/json", bytes.NewBuffer(postData))
	if postErr != nil {
		return postErr
	}
	defer resp.Body.Close()

	body, loadErr := ioutil.ReadAll(resp.Body)
	if loadErr != nil {
		return loadErr
	}
	return json.Unmarshal(body, &a.responseType)
}

// Get post
// おそらくリクエストパラメータの設定をどうするか色々出てくると思うので、
// いったん仮置で
func (a *APIClient) Get() error {
	resp, postErr := http.Get(a.url)
	if postErr != nil {
		return postErr
	}
	defer resp.Body.Close()

	body, loadErr := ioutil.ReadAll(resp.Body)
	if loadErr != nil {
		return loadErr
	}
	return json.Unmarshal(body, &a.responseType)
}
