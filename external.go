package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "bytes"
    "log"
)

type LinkedInCall struct {
  Comment string
  Content map[string]string
  Visibility map[string]string
}

type LinkedInResponse struct {
  UpdateKey, UpdateURL string
}

func RequestLinkedIn(m *Message) ([]byte, LinkedInResponse) {
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
  // equivalent to Marshal, except intended for streams
  // NewEncoder accepts an io.Writer, so it's generally intended for streams
  // Encode writes the JSON encoding of the argument passed to that stream
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
  // defer moves this to the end of the function


  var lrs LinkedInResponse
  fmt.Println("response Status: ", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal(body, &lrs)
  // decode body out of JSON into &lrs
  fmt.Println("response Body:", string(body))
  fmt.Println(lrs)
  return body, lrs;
}
