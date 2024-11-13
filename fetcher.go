package main

import (
  "crypto/tls"
  "github.com/tidwall/gjson"
  "io/ioutil"
  "net/http"
  "net/url"
)

const API_KEY = "XYZ"

func fetchJSON() (string, error) {
  u := url.URL{
    Scheme: "https",
    Host: "192.168.0.1:3000",
    Path: "/lua/rest/v2/get/interface/data.lua",
  }

  p := u.Query()
  p.Set("ifid", "0")
  u.RawQuery = p.Encode()
  req, err := http.NewRequest("GET", u.String(), nil)
  if err != nil {
    return "", err
  }
  req.Header.Add("Authorization", "Token " + API_KEY)
  client := &http.Client{
    Transport: &http.Transport{
      TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    },
  }
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func fetchUpDown() (float64, float64, error) {
  json, err := fetchJSON()
  if err != nil {
    return 0, 0, err
  }
  up := gjson.Get(json, "rsp.troughput.upload.bps").Float()
  down := gjson.Get(json, "rsp.troughput.download.bps").Float()
  return down / 1000, up / 1000, nil
}
