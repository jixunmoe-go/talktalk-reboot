package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type payloadBase struct {
	CSRF csrfParams `json:"csrf"`
}

func (r *Client) url(path string) string {
	return r.BaseURL + path
}

func (r *Client) post(path string, data, result interface{}) {
	body, err := json.Marshal(data)
	check(err)

	resp, err := r.client.Post(r.url(path), "application/json", bytes.NewReader(body))
	check(err)

	err = parseResp(resp.Body, result)
	check(err)
}

func parseResp(stream io.ReadCloser, out interface{}) error {
	body, err := ioutil.ReadAll(stream)
	if err != nil {
		return fmt.Errorf("could not read stream: %w", err)
	}

	// some response are surrounded with "while(0); /*" and "*/"
	if body[0] != '{' {
		body = body[12 : len(body)-2]
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return fmt.Errorf("could not parse json: %w", err)
	}

	return nil
}
