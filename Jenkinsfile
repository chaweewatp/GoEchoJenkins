pipeline {
agent any
tools {
    go 'go-1.11'
}
environment {
    GO111MODULE = 'on'
}
stages {


stage('Compile') {
        steps {
//          sh 'ps -ef | grep main10 | grep -v grep | awk "{print $2}" | xargs kill'
sh '''
if lsof -i:1323
then
kill $(lsof -t -i:1323)
fi
'''
            sh 'go mod init'
            sh 'go get github.com/labstack/echo/v4'
            sh 'go build'
        }
    }

stage('Release'){
steps{
  sh 'go run main10.go'
}
}

//    stage('Test') {
//            environment {
//                CODECOV_TOKEN = credentials('codecov_token')
//            }
//            steps {
//                sh 'go test ./... -coverprofile=coverage.txt'
//                sh "curl -s https://codecov.io/bash | bash -s -"
//            }
//        }
//        stage('Code Analysis') {
//            steps {
//                sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
//                sh 'golangci-lint run'
//            }
//        }
//        stage('Release') {
//            when {
//                buildingTag()
//            }
//            environment {
//                GITHUB_TOKEN = credentials('github_token')
//            }
//            steps {
//                sh 'curl -sL https://git.io/goreleaser | bash'
//            }
//        }
}

}
