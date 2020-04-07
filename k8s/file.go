package k8s


func CreateKubeconfigFile(user,server,clustername string)string{
  return `apiVersion: v1
kind: Config
clusters:
- cluster:
    certificate-authority-data: `+getCertificateAuthData(clustername)+`
    server: `+server+`
  name: `+clustername+`
users:
- name: `+user+`
  user:
    client-certificate-data: `+getClientCertificateData(user)+`
contexts:
- context:
    cluster: `+clustername+`
    user: `+user+`
  name: `+user+`-`+clustername+`
current-context: `+user+`-`+clustername
}
