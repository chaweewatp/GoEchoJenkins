
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
//         stage('deploy'){
//             steps {
//                 sh '''
//                     ./main10
//                 '''
//             }
//         }


        stage('Setup gunicorn service'){
            steps {
                sh '''
                    chmod +x gunicorn.sh
                    ./gunicorn.sh
                    '''
            }
        }
//         stage('Setup Nginx'){
//             steps {
//                 sh '''
//                     chmod +x nginx.sh
//                     ./nginx.sh
//                     '''
//             }
//         }
    }
}
