-- This is lua module. Used in example.lua
local mod = {}

local kubernetes = require("kubernetes")

mod.deployment = kubernetes.Deployment()

mod.deployment.metadata = {
  name = "deployment",
  namespace = "default",
  labels = {
    name = "deployment"
  }
}


mod.deployment.spec = {
  replicas = 4,
  selector = {
    matchLabels = "name=deployment"
  },

  containers = {
    {
      image = "nginx:nginx",
      name = "container-name"
    }
  }
}
return mod