pipeline {
  agent any
  stages {
    stage('Checkout Code') {
      steps {
        git(url: 'https://github.com/codelamps-academy/belajar-golang', branch: 'development')
      }
    }

    stage('') {
      steps {
        sh 'ls -la'
      }
    }

  }
}