package k8s

import (
  "os"
  "fmt"
  "log"
  "bytes"
  "encoding/json"
  "io/ioutil"
  "os/exec"
)




func KubectlApply(ressource string){
  os.Create("tmp.yaml")
  err := ioutil.WriteFile("tmp.yaml", []byte(ressource), 0644)
  cmd := exec.Command("kubectl","apply","-f","tmp.yaml")
  var out bytes.Buffer
  cmd.Stdout = &out
  err = cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  os.Remove("tmp.yaml")
  fmt.Println(out.String())
}


func ApproveCert(user string){
  cmd := exec.Command("kubectl","certificate","approve",user+"-csr")
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
}



func getClientCertificateData(user string)string{
  cmd := exec.Command("kubectl","get","csr",user+"-csr","-o","jsonpath='{.status.certificate}'")
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  return out.String()
}

func getCertificateAuthData(clustername string)string{
  cmd := exec.Command("kubectl","config","view","--raw","-o","json")
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  var a interface{}
  err = json.Unmarshal( []byte(out.String()), &a)
  if err != nil {
    log.Fatal(err)
  }
  b := a.(map[string]interface{})["clusters"].([]interface{})
  for _,c := range b {
    d := c.(map[string]interface{})
    if d["name"] == clustername {
      return d["cluster"].(map[string]interface{})["certificate-authority-data"].(string)
    }
  }
  return ""
}
