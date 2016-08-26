package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "bytes"
    "log"
)

type Message struct {
  MessageURL, Token, Platform string
}

type LinkedInCall struct {
  Comment string
  Content map[string]string
  Visibility map[string]string
}

type FacebookCall struct {
  Message string `json:"message"`
  Access_token string `json:"access_token"`
}

type LinkedInResponse struct {
  UpdateKey, UpdateURL string
}

type FacebookResponse struct {
  Id string
}

func (m *Message) LinkedIn() ([]byte, LinkedInResponse) {
  cont := map[string]string{
    "title": "Hack Reactor and You",
    "description": "pillows",
    "submitted-url": m.MessageURL,
    "submitted-image-url": "https://golang.org/doc/gopher/frontpage.png",
  }

  vis := map[string]string{
    "code": "anyone",
  }

  bod := LinkedInCall{
    "Pillow Talk",
    cont,
    vis,
  }

  b := new(bytes.Buffer)
  json.NewEncoder(b).Encode(bod)
  url := `https://api.linkedin.com/v1/people/~/shares?oauth2_access_token=`+m.Token+`&format=json`
  req, err := http.NewRequest("POST", url, b)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("x-li-format", "json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err.Error())
  }
  defer resp.Body.Close()

  var lrs LinkedInResponse
  fmt.Println("response Status: ", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal(body, &lrs)
  fmt.Println("response Body:", string(body))
  fmt.Println(lrs)
  return body, lrs;

}

func (m *Message) Facebook() ([]byte, FacebookResponse) {
  fbjson := FacebookCall{m.MessageURL, m.Token}
  b := new(bytes.Buffer)
  json.NewEncoder(b).Encode(fbjson)
  fmt.Println(b.String());
  url := `https://graph.facebook.com/me/feed`
  req, err := http.NewRequest("POST", url, b)
  if err != nil {
    fmt.Println(err)
  }
  req.Header.Set("Content-Type", "application/json")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()

  var frs FacebookResponse
  fmt.Println("response Status: ", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal(body, &frs)
  fmt.Println("response Body:", string(body))
  fmt.Println(frs)
  return body, frs;
}
