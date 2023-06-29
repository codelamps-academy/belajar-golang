pipeline {
  agent any
  stages {
    stage('Checkout Code') {
      steps {
        git(url: 'https://github.com/codelamps-academy/belajar-golang', branch: 'development')
      }
    }

    stage('Show Shell') {
      steps {
        sh 'ls -la'
      }
    }

    stage('Print') {
      steps {
        echo 'Hai, you using jenkins'
      }
    }

    stage('Build') {
      steps {
        sh 'docker build -f Dockerfile .'
      }
    }

    stage('Login into Dockerhub') {
      environment {
        DOCKERHUB_USER = 'andre'
        DOCKERHUB_PASSWORD = 'andre'
      }
      steps {
        sh 'docker login -u $DOCKERHUB_USER -p $DOCKERHUB_PASSWORD'
      }
    }

  }
}