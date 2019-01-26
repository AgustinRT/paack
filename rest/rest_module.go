package rest

import (
  	"net/http"
    "bytes"
    "log"
  )

// url with a fake server maked with node for emulate the quiz behaviour
const (
	url = "http://localhost:3000/api/v1/person"
)

func Do_post(dataBytes string) (int, string) {
  resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(dataBytes)))
  if err != nil {
    log.Println(err)
  }
  buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  return resp.StatusCode, buf.String()
}

// func Do_get()  {
//
// }
