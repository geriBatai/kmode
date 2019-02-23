# KMode

This is a very early prototype for a Kubernetes deployment tool. Currently it just prints out generated objects to the
standard output, but the intention is to build a proper deployment tool, with diffs, deployments, modules and resource
dependencies. Additional and independent terraform provider should be trivial to implement.

## Getting started

Look at `example/example.lua` for the way resources will be defined. Currently only few
objects are implemented, but others are rather trivial to implement. To build it with any recent enough go version (1.11):

```
go build
```

To run example code:

```
kmode output --var-file example/vars/production.lua --filename example/example.lua
```

kthxbye.