pipeline {
   agent {
      kubernetes {
         yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    agent-type: golang
spec:
  containers:
  - name: golang
    image: golang:1.16.3-alpine
    command:
    - cat
    tty: true
"""
      }
   }

   stages {
      stage('Hello') {
         steps {
            echo "Hello World"
            container('golang') {
               sh """
                go run hello-world.go
                go build hello-world.go
                ./hello-world
                """
            }
         }
      }
   }
}
