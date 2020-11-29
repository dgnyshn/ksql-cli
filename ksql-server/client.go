package ksql_server

import (
	env "../initCfg"
	"io"
	"io/ioutil"
	"net/http"
)

func NewRequest(method string, url string, request io.Reader) []byte {
	req, err := http.NewRequest(method, url, request)

	if err != nil {
		panic(err)
	}

	if env.DefaultAuthentication != "" {
		req.Header.Set("Authentication", env.DefaultAuthentication)
	}
	req.Header.Set("Accept", "application/vnd.ksql.v1+json")

	client := &http.Client{}
	res, err := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body
}
