package main

import (
  "github.com/urfave/cli"
  "log"
  "os"
  "io/ioutil"
  k8s "./k8s"
)

func main() {
    var username,namespace,server,clustername,csrpath string
    app := &cli.App{
    Name: "usercert",
    Usage: "Create certification log on for a specific user",
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name: "username",
        Value: "default",
        Usage: "Name of the created user",
        Destination: &username,
      },
      &cli.StringFlag{
        Name: "namespace",
        Value: "default-ns",
        Usage: "Namespace where the user can act",
        Destination: &namespace,
      },
      &cli.StringFlag{
        Name: "server",
        Value: "http://localhost:8080",
        Usage: "url to the k8s server",
        Destination: &server,
      },
      &cli.StringFlag{
        Name: "clustername",
        Value: "mycluster",
        Usage: "Name of the cluser will be use in the kubeconfig file)",
        Destination: &clustername,
      },
      &cli.StringFlag{
        Name: "csr",
        Value: "./mycsr.csr",
        Usage: "Path to the csr file",
        Destination: &csrpath,
      },
    },
    Action: func(c *cli.Context) error {
      k8s.KubectlApply(k8s.CreateNamespace(namespace))
      k8s.KubectlApply(k8s.CreateRole(username,namespace))
      k8s.KubectlApply(k8s.CreateRoleBinding(username,namespace))
      k8s.KubectlApply(k8s.CreateCSR(username,csrpath))
      k8s.ApproveCert(username)
      os.Create(clustername+"kubeconfig")
      err := ioutil.WriteFile(username+"-kubeconfig", []byte(k8s.CreateKubeconfigFile(username,server,clustername)), 0644)
      return err
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
