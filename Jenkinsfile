pipeline {
  agent none
  stages {
    stage('Build and Test') {
      agent {
        docker {
          image 'golang:latest'
          args '-e GOCACHE=/home/ubuntu/jenkins_node/workspace/pipeline_lesson_9'
        }
      }
      stages {
        stage('Build') {
          steps {
            echo 'Building application...'
            sh 'echo $GOCACHE'
            sh 'go build ./web_app.go'
          }
        }
        stage('Test') {
          steps {
            echo 'Testing application...'
            sh './web_app &'
            sh '''response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8081) &&
            echo Resoince is: $response &&
            [ "$response" = "200" ] && exit 0 || exit 1'''
          }
        }
      }
    }
    stage('Deploy') {
      agent { label 'node1' }
      steps {
        echo 'Deploying application...'
        sshagent(credentials: ['ubuntu-creds']) {
            sh 'ssh -o StrictHostKeyChecking=no ubuntu@3.77.236.209 "cd /home/ubuntu/web_app && sudo git pull && sudo go build ./web_app.go && sudo systemctl restart web_app"'
        }
      }
    }
  }
}
