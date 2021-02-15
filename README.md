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
