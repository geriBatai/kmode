local kubernetes = require("kubernetes")

-- print(w)


deployment = kubernetes.Deployment()

deployment.metadata = {
  name = "deployment",
  namespace = "default",
  labels = {
    name = "deployment"
  }
}

deployment.spec = {
  replicas = 10,
  selector = {
    matchLabels = {
      "name=deployment"
    }
  },

  template = {
    spec = {
      containers = {
        {
          image = "nginx:nginx",
          name = "container-name"
        }
      }
    }
  }
}