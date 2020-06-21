pipeline {
  agent any
  stages {
    stage('unittest') {
      steps {
        sh 'echo unitest'
      }
    }

    stage('build') {
      steps {
        sh 'echo docker build'
      }
    }

    stage('deploy_to_test') {
      steps {
        sh 'echo deploy_to_test'
      }
    }

    stage('intergration_test') {
      steps {
        sh 'echo intergration test'
      }
    }

    stage('update_installer') {
      steps {
        sh 'echo update installer'
        sh 'git clone git@github.com:maplefeng-a/hello_go.git'
      }
    }

    stage('push_to_harbor') {
      steps {
        sh 'echo push to harbor'
      }
    }

    stage('deploy') {
      steps {
        sh 'echo deploy'
      }
    }

  }
  environment {
    testnode_ip = '192.168.5.105'
  }
}
