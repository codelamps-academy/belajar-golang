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

  }
}