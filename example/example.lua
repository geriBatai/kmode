
-- First we need to import our kubernetes module
local kubernetes = require("kubernetes")

-- That's how we define variables. We can assign this
-- to the object later. You can define environment
-- specific variables in vars/ files.
local metadata = {
  name = "hi",
  namespace = "default"
}

-- Let's create service and assign metadata
-- from the above. Then, change the name in a 
-- map to something else.
svc = kubernetes.Service()
svc.metadata = metadata
svc.metadata.name = "some-service"

-- Let's define local deployment. Local variables
-- do not generate output.
local deployment = kubernetes.Deployment()
deployment.metadata = metadata

-- Create two identical copies for deployment
-- object. As these are global variables, they
-- will be present in generated output
deployment2 = deployment:clone()
deployment3 = deployment:clone()

-- We could use setNamespace method instead
-- of assigning deployment3.metadata.namespace
deployment3:setNamespace("custom-namespace")

-- Set names for deployment2 and deployment3.
-- These two methods are identical
deployment2.metadata.name = "deployment2"
deployment3:setName("deployment3")

-- As this is lua, we have conditionals
-- and other stuff you could expect.
if environment == "production" then
  deployment2:setNamespace("prod")
  deployment2.metadata.labels = labels
  deployment2.spec.selector.matchLabels = labels
end

-- We can even use lua modules. In this case,
-- load mod.lua
local mod = require("example/mod")
deployment4 = mod.deployment:Clone()

-- output is generated once all of this code is executed
