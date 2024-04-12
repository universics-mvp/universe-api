APPNAME = 'appname'

branch_name = BRANCH_NAME
PULL_REQUEST = false
try {
  branch_name = CHANGE_BRANCH
  echo "It's pull request"
  PULL_REQUEST = true
} catch (Exception ex) {}

podTemplate(containers: [
    containerTemplate(name: 'docker', image: 'docker:dind', ttyEnabled: true, command: 'cat', privileged: true),
    containerTemplate(name: "golang", image: "klutrem/golang-java:1.1", command: "sleep", args: "99d", NodeSelector: "kubernetes.io/os=linux")
  ],
  volumes: [
    hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock')
  ], hostNetwork: true)  {
    node(POD_LABEL) {
    if (BRANCH_NAME.contains('PR')) {
        BRANCH = CHANGE_BRANCH
    } else {
        BRANCH = BRANCH_NAME
    }
    if (env.BRANCH_NAME == "main") {
      ENV = "prod"
    } else {
      ENV = env.BRANCH_NAME
    }

    helpersPath = "/tmp/helpers.groovy"
			sh "curl -L ${env.GIT_SCHEME}${env.GIT_HOST}/Jenkins/jenkins-tools/raw/branch/master/helpers.groovy -o $helpersPath"
			helpers = load helpersPath

      /* -------------------------------------------------------------------------- */
      /*                                    pull                                    */
      /* -------------------------------------------------------------------------- */
      stage('pull') {
        helpers.withGiteaCreds("GitBackendCreds", "git-tool") {
          REPO_URL = "${env.GIT_SCHEME}${env.GIT_HOST}/${env.GITEA_KAISER_ORG}/${APPNAME}.git"
          echo "Pulling $REPO_URL[$BRANCH_NAME]..."
          sh "git clone $REPO_URL --branch $BRANCH ."
          sh "git submodule update --init --recursive"
        }
      }
      container("golang") {
        /* -------------------------------------------------------------------------- */
        /*                                    build                                   */
        /* -------------------------------------------------------------------------- */
        stage("build") {
          sh "go get ./..."
        }
        /* -------------------------------------------------------------------------- */
        /*                                    test                                    */
        /* -------------------------------------------------------------------------- */
        stage("test") {
          sh "go test ./... -coverprofile=coverage.out"
        }
      }
      /* -------------------------------------------------------------------------- */
      /*                            skip if pull request                            */
      /* -------------------------------------------------------------------------- */
      if(!PULL_REQUEST) {
        /* -------------------------------------------------------------------------- */
        /*                                 build image                                */
        /* -------------------------------------------------------------------------- */
        stage('build docker image') {
          container('docker') {
            IMG = "${env.DOCKER_REGISTRY_HOST}/${env.DOCKER_REGISTRY_ORGANISATION_NAME}/${APPNAME}_${ENV}:${BUILD_NUMBER}"
            echo "Pushing $IMG to docker registry ${env.DOCKER_REGISTRY_HOST}"
            sh "docker build -t ${IMG} ."
          }
        }
        /* -------------------------------------------------------------------------- */
        /*                                 push image                                 */
        /* -------------------------------------------------------------------------- */
        stage("push image") {
          withCredentials([
            usernamePassword(
              credentialsId: 'DockerRegistryCreds',
              usernameVariable: 'DOCKER_LOGIN',
              passwordVariable: 'DOCKER_PASSWORD',
            )
          ]) 
          {
            container('docker') {
              sh "docker login ${env.DOCKER_REGISTRY_HOST} -u $DOCKER_LOGIN -p $DOCKER_PASSWORD"
              sh "docker push $IMG"
            }
          }
        }
        /* -------------------------------------------------------------------------- */
        /*                                  sonarqube                                 */
        /* -------------------------------------------------------------------------- */
        stage('SonarQube Analysis') {
          container("golang") {
            def scannerHome = tool 'main';
            withSonarQubeEnv() {
              sh "${scannerHome}/bin/sonar-scanner"
              sh "rm -rf coverage"
            }
          }
        }
        /* -------------------------------------------------------------------------- */
        /*                                   cleanup                                  */
        /* -------------------------------------------------------------------------- */
        stage ('cleanup'){
          cleanWs()
        }
      }
    }
  }
  