
    node{
    // Install the desired Go version
    def root = tool name: 'Go-Tool', type: 'go'
    stages {
        stage ('Checkout Stage') {

            steps {
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    git "https://github.com/ghoshTubai/firstgolang.git"
                }
            }
        }
        stage ('Testing Stage') {
            steps {
               withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'go test .'
                }
            }
        }


        stage ('Deployment Stage') {
            steps {
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'go install .'
                }
            }
        }
    }
    }
