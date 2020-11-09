pipeline {
    agent any
    tools {
        go 'golang'
        dockerTool 'docker'
    }
    stages {
        stage ('Installing dependencies'){
            steps {
                echo 'Installing dependencies'
                sh 'go get -u  github.com/go-sql-driver/mysql'                
            }
            
        }

        stage ('Git'){
            steps {
                echo 'Getting Git'
                git url: 'https://github.com/entrust-albert/registration_service'
            }
        }
        
        stage ('Building'){
            steps {
                echo 'Compiling and building'
                sh 'go build -o poster main.go'
            }
            
        }

        stage('Dockerize') {
            steps{
                sh 'docker build -t post:v1.0 .'
            }
            
        }
        
    }
    
}
