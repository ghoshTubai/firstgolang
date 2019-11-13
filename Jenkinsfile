pipeline {
    agent any
    tools {
        go 'Go-Tool'
    }
    environment {
        GO143MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}