package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://gateway.xoniaapp.com/api/account/login"
  method := "POST"

  payload := strings.NewReader(`{
    "email":"demo@demo.com",
    "password": "password"
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Cookie", "xa=MTY0NzU5ODg1OXxOd3dBTkZkRlFrVlVVVUpUTjFaTVIwSXpOVFJMVEV4VE16WlRWVUpOUjBGQ1Z6Uk5OVGRNVkVKTlJrNVpVRWMzVkRSSlExWlBObEU9fB029RurEhFfANCHjHJLTS_ShGUcH3C5UI632L2o-2FL")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
