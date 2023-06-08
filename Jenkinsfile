pipeline {
  agent none
  stages {
    stage('Build and Test') {
      agent {
        docker {
          image 'golang:latest'
          args '-e GOCACHE=/home/ubuntu/jenkins_node/workspace/pipeline_lesson_10'
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
        stage('Archive') {
          steps {
            echo 'Archiving artifacts...'
            archiveArtifacts artifacts: '**/web_app', allowEmptyArchive: false, followSymlinks: false
          }
        }
      }
    }
    stage('Deploy') {
      agent { label 'node1' }
      steps {
        echo 'Deploying application...'
        sshagent(credentials: ['ubuntu-creds']) {
          sh "ssh -o StrictHostKeyChecking=no ubuntu@3.71.186.91 'cd web_app && sudo systemctl stop web_app && rm -f web_app && wget -O web_app https://mbps54-bucket-artif-1.s3.eu-central-1.amazonaws.com/pipeline_lesson_10/${env.BUILD_NUMBER}/artifacts/web_app && sudo chmod +x web_app && sudo systemctl start web_app'"
        }
      }
    }
  }
}
