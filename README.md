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

# Testing steps

- Log in to the VM
- Make sure that all the pods are running

```
kubectl  -n icap-adaptation get pods
```

- Start a test using the command bellow : If all is ok you will receive a result file.

```
mkdir /tmp/input
cp <pdf_file_name> /tmp/input/
docker run --rm -v /tmp/input:/opt/input -v /tmp/output:/opt/output glasswallsolutions/c-icap-client:manual-v1 -s 'gw_rebuild' -i <your vm IP> -f '/opt/input/<pdf_file_name>' -o /opt/output/<pdf_file_name> -v
```

During the test review the pods logs (icap-server, adaptation-service, any rebuild pods)

# Rebuild flow to implement

![new-rebuild-flow-v2](https://user-images.githubusercontent.com/76431508/107766490-35064200-6d3c-11eb-8d63-ad64f29ce964.jpeg)
