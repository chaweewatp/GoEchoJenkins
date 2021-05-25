pipeline {
    agent any
    tools {
        go 'go-1.16'
    }
    environment {
        GO111MODULE = 'on'
    }
    stage('Compile') {
        steps {
            sh 'go build'
        }
    }

}
