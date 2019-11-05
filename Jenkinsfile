def imageName = 'mlabouardy/imdb-engine'

node('slaves'){
    stage('Checkout'){
        checkout scm
    }

    stage('Quality Test'){
        docker.build("${imageName}:${env.BUILD_ID}", '-f Dockerfile.quality .')
        sh "docker run --rm ${imageName}:${env.BUILD_ID}"
    }

    stage('Unit Test'){
        docker.build("${imageName}:${env.BUILD_ID}", '-f Dockerfile.unit .')
        sh "docker run --rm ${imageName}:${env.BUILD_ID}"
    }

    stage('Security Test'){
        docker.build("${imageName}:${env.BUILD_ID}", '-f Dockerfile.security .')
        sh "docker run --rm ${imageName}:${env.BUILD_ID}"
    }

    stage('Build'){
        docker.build(imageName)
    }

    /*stage('Push'){
        docker.withRegistry(registry, 'registry') {
            docker.image(imageName).push("${commitID()}")

            if (env.BRANCH_NAME == 'master') {
              docker.image(imageName).push('latest')
            }
            if (env.BRANCH_NAME == 'preprod') {
              docker.image(imageName).push('preprod')
            }
            if (env.BRANCH_NAME == 'develop') {
              docker.image(imageName).push('develop')
            }
        }
    }

    stage('Deploy'){
        build job: "dlf-deployment/master"
    }*/
}

def commitID() {
    sh 'git rev-parse HEAD > .git/commitID'
    def commitID = readFile('.git/commitID').trim()
    sh 'rm .git/commitID'
    commitID
}