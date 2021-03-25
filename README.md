<h1 align="center">go-k8s-controller</h1>

<p align="center">
    <a href="https://github.com/k8-proxy/go-k8s-controller/actions/workflows/build.yml">
        <img src="https://github.com/k8-proxy/go-k8s-controller/actions/workflows/build.yml/badge.svg"/>
    </a>
    <a href="https://codecov.io/gh/k8-proxy/go-k8s-controller">
        <img src="https://codecov.io/gh/k8-proxy/go-k8s-controller/branch/main/graph/badge.svg"/>
    </a>	    
    <a href="https://goreportcard.com/report/github.com/k8-proxy/go-k8s-controller">
      <img src="https://goreportcard.com/badge/k8-proxy/go-k8s-controller" alt="Go Report Card">
    </a>
	<a href="https://github.com/k8-proxy/go-k8s-controller/pulls">
        <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome">
    </a>
    <a href="https://opensource.org/licenses/Apache-2.0">
        <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="Apache License, Version 2.0">
    </a>
    <a href="https://github.com/k8-proxy/go-k8s-controller/releases/latest">
        <img src="https://img.shields.io/github/release/k8-proxy/go-k8s-controller.svg?style=flat"/>
    </a>
</p>

# Rebuild pod controller

This service is a controller service making sure we always have a certain number of hot pods running and listening to the queues. Ready to get a file for processing.

### Steps of processing
When it starts
- It will start an amount of pods
- It will monitor their status and in case one of them get completed it triggers a new one

## Configuration
These environment variables are needed by the service 
- POD_COUNT : Count of pods to start
- MINIO_USER : Minio access key
- MINIO_PASSWORD : Minio access secret


### Docker build
- To build the docker image
```
docker build -t <docker_image_name> .
```

- This works only on a kubernetes cluster, deploiement steps available on 

# Testing steps

- Log in to the VM
- Make sure that all the pods are running

```
kubectl  -n icap-adaptation get pods
```

- To test, just try to rebuild a file, a rebuild pod will pick it up and a new one should be created

During the test review the pods logs (icap-server, adaptation-service, any rebuild pods)

# Rebuild flow to implement
![new-rebuild-flow-v2](https://github.com/k8-proxy/go-k8s-infra/raw/main/diagram/go-k8s-infra.png)
