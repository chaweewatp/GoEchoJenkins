
pipeline {
    agent any
    stages {
        stage('init'){
        steps{
        sh '''
        export PATH="/usr/local/go/bin/:$PATH"
        go version
        go mod init myapp
        go get github.com/labstack/echo/v4
        '''
        }
      }
         stage('build'){
            steps {
                sh '''
                  go build main10.go
                    '''
            }
        }
         stage('deploy'){
             steps {
                 sh '''
                     go run main10.go
                 '''
             }
         }
    }
}
