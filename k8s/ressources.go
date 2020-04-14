package k8s

import (
  utils "../utils"
)


func CreateCSR(user,csrpath string)string{
  return `apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: `+user+`-csr
spec:
  groups:
  - system:authenticated
  request: `+utils.EncodeCSR(csrpath)+`
  usages:
  - digital signature
  - key encipherment
  - server auth
  - client auth`
}

func CreateNamespace(ns string)string{
  return `---
apiVersion: v1
kind: Namespace
metadata:
  name: `+ns
}

func CreateRole(user,ns string)string{
  return `kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: `+ns+`
  name: `+user+`-role
rules:
- apiGroups: [""]
  resources: ["pods", "services"]
  verbs: ["create", "get", "update", "list", "delete", "patch"]
- apiGroups: ["apps"]
  resources: ["deployments","daemonset"]
  verbs: ["create", "get", "update", "list", "delete", "patch"]`
}

func CreateRoleBinding(user,ns string)string{
  return `kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: `+user+`-rb
  namespace: `+ns+`
subjects:
- kind: User
  name: `+user+`
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: `+user+`-role
  apiGroup: rbac.authorization.k8s.io`
}
