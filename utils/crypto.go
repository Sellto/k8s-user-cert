package utils

import (
  "encoding/base64"
  "io/ioutil"
)

func EncodeCSR(path string) string {
  dat, err := ioutil.ReadFile(path)
  if err != nil {
        panic(err)
  }
  return base64.StdEncoding.EncodeToString(dat)
}
