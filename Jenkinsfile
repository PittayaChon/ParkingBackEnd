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
                sh 'docker image rm parkingbackend'
                sh 'docker image prune -a'
                sh 'docker volume prune'
                sh 'docker build -t parkingbackend .'
                //   script{
                // //    app = docker.build('api:latest')
                //    dockerImage = docker.build registry + ":latest"
                //   }
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
