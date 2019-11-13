node {
  checkout scm

  docker.withRegistry("https://registry.hub.docker.com","")
        def customImage = docker.build("ankur/go")
        customImage.push()
}