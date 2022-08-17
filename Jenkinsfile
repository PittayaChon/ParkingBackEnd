pipeline {
    environment {
    registryCredential = ''
    dockerImage = ''
  }
  agent  any
    stages {
        stage('Clone git') {
            steps {
                script{
                checkout scm
                }
            }
        }

        stage('Build image') {
            steps {
                  script{
                  dockerImage = docker.build 'api:latest'
                //   apiImage = docker.build('0865079783/api:latest')
                  }
            }
        }

        stage('Testing') {
            steps {
                echo 'Testing..'
            }
        }

        stage('Push image') {
            steps {
                  echo '..'
            }
        }

        stage('Prepare deploy') {
            steps {
                    echo 'Building..'
            }
        }

        stage('Deploy on production') {
            steps {
              script {
                    echo 'Building..'
                }
              }
        }
    }
}
