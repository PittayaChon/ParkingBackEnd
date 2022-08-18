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
                sh 'docker image rm parkingbackend -f'
                sh 'docker image prune -a -f'
                sh 'docker volume prune -f'
                sh 'docker build -t 0865079783/parkingbackend .'
            }
        }

        stage('Testing') {
            steps {
                echo 'Testing..'
            }
        }

        stage('Push image') {
            steps {
                  sh 'docker login -u="0865079783" -p="dearx2527"'
                  sh 'docker push 0865079783/parkingbackend:latest'
            }
        }

        stage('Prepare deploy') {
            steps {
                    sshagent(credentials: ['jenkins-production']) {
                    sh 'ssh -o StrictHostKeyChecking=no ubuntu@prod.sandbox-me.com mkdir -p /home/ubuntu/parkingfontend'
                    sh 'scp -o StrictHostKeyChecking=no docker-compose.yml prod.sandbox-me.com:/home/ubuntu/parkingfontend/docker-compose.yml'
                    // sh 'scp -o StrictHostKeyChecking=no .env ${production_server_ip}:/home/dong/todoapp/.env'
                    // scp -i  docker-compose.yml ubuntu@prod.sandbox-me.com:/parkingfontend
                    }
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
