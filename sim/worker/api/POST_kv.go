package api

import (
  "bytes"
  "encoding/json"
  "net/http"
)

type POST_kv_request{
  Consistency string `json:"consistency"`
}

type POST_kv_response{
  Code int `json:"code"`
  Id string `json:"id"`
}

func (self *API)POST_kv()(int, string, error){
  request := POST_kv_request{Consistency:"none"}
  json_buf, err := json.Marshal(request)
  body := bytes.NewBuffer(json_buf)
  req, err := http.NewRequest("GET", self.URL + "/kv", body)
  if err != nil{
    return -1, "", err
  }
  req.Header.Add("content-type", "application/json; charset=utf-8")

  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil{
    return -1, "", err
  }

  resBody, err := io.ReadAll(res.Body)
  if err != nil{
    return -1, "", err
  }
  defer res.Body.Close()

  var ctx POST_kv_response
  err = json.Unmarshal(resBody, &ctx)
  if err != nil{
    return -1, "", err
  }

  self.Kv_id = ctx.Id
  return ctx.Code, string(resBody), err
}
