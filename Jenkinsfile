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
        stage('Checkout SCM') {
          steps {
            checkout([
              $class: 'GitSCM',
              branches: [[name: 'main']],
              userRemoteConfigs: [[
                url: 'https://github.com/mbps54/web_app.git',
                credentialsId: '',
              ]]
            ])
          }
        }
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
        sshagent(credentials: ['ubuntu']) {
            sh 'ssh -o StrictHostKeyChecking=no 3.120.40.38 "cd /home/ubuntu/web_app && git pull && go build ./web_app.go && sudo systemctl restart web_app"'
        }
      }
    }
  }
}
