package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GET_kv_id_field_key_request struct {
	NameSpace string `json:"id"`
}

type GET_kv_id_field_key_response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Kvid   string `json:"id"`
}

func (self *API) GET_kv_id_field_key(key string) (int, string, error) {
	request := GET_kv_id_field_key_request{self.Kv_id}
	json_buf, err := json.Marshal(request)
	if err != nil {
		return -1, "", err
	}

	body := bytes.NewBuffer(json_buf)
	req, err := http.NewRequest("GET", self.URL+"/kv/", self.Kv_id+"/field/"+key)
	if err != nil {
		return -1, "", err
	}
	req.Header.Add("content-type", "application/json; charset=utf-8")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
}
