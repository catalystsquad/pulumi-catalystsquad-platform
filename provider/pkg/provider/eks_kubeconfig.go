package provider

// this thing is so ugly, put it in it's own file
var eksDefaultKubeConfig string = `apiVersion: v1
clusters:
- cluster:
    server: %s
    certificate-authority-data: %s
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: aws
  name: aws
current-context: aws
kind: Config
preferences: {}
users:
- name: aws
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: aws
      args:
      - "eks"
      - "get-token"
      - "--cluster-name"
      - "%s"
`

var kubeConfigArgsRoleArn string = `      - "--role-arn"
      - "%s"
`

var kubeConfigArgsAwsProfile string = `      env:
      - name: AWS_PROFILE
        value: "%s"
`
