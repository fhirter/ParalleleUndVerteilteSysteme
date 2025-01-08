# Lösung Übung Microservices

### minikube

The official minikube documentation is pretty good: 
[minikube.sigs.k8s.io/docs/start](https://minikube.sigs.k8s.io/docs/start/?arch=%2Fmacos%2Fx86-64%2Fstable%2Fbinary+download)

- `brew install minikube`
- `brew install kubectl`
- Show Dashboard: `minikube start`
- List all pods: `kubectl get pods -A`
```
NAMESPACE              NAME                                        READY   STATUS             RESTARTS      AGE
kube-system            coredns-6f6b679f8f-t6dx2                    1/1     Running            1 (21m ago)   5h38m
kube-system            coredns-6f6b679f8f-vt9ht                    1/1     Running            1 (21m ago)   5h38m
kube-system            etcd-minikube                               1/1     Running            1 (21m ago)   5h38m
kube-system            kube-apiserver-minikube                     1/1     Running            1 (20m ago)   5h38m
kube-system            kube-controller-manager-minikube            1/1     Running            1 (21m ago)   5h38m
kube-system            kube-proxy-4ktld                            1/1     Running            1 (21m ago)   5h38m
kube-system            kube-scheduler-minikube                     1/1     Running            1 (21m ago)   5h38m
kube-system            storage-provisioner                         1/1     Running            4 (33s ago)   5h38m
kubernetes-dashboard   dashboard-metrics-scraper-c5db448b4-k5fbk   1/1     Running            1 (21m ago)   5h33m
kubernetes-dashboard   kubernetes-dashboard-695b96c756-fxwhg       1/1     Running            2 (33s ago)   5h33m
```
- Install Minikube Dashboard: `minikube dashboard`
- build and push docker image:
```bash
set -o errexit   # abort on nonzero exitstatus
set -o nounset   # abort on unbound variable
set -o pipefail  # don't hide errors within pipes
cd node
image_name=fhirter/node-server
docker build -t "$image_name" .
docker push "$image_name"
```
- create deployment: `kubectl create deployment node-server --image=fhirter/node-server:latest`
- `kubectl get pods`, status should be `Running`
```
NAME                         READY   STATUS    RESTARTS   AGE
node-server-ddbd9b95-g8h5s   1/1     Running   0          30s
```
- expose service: `kubectl expose deployment node-server --type=NodePort --port=8080`
- Connect to service: `minikube service node-server`
```
|-----------|-------------|-------------|---------------------------|
| NAMESPACE |    NAME     | TARGET PORT |            URL            |
|-----------|-------------|-------------|---------------------------|
| default   | node-server |        8080 | http://192.168.49.2:30569 |
|-----------|-------------|-------------|---------------------------|
🏃  Start Tunnel für den Service node-server
|-----------|-------------|-------------|------------------------|
| NAMESPACE |    NAME     | TARGET PORT |          URL           |
|-----------|-------------|-------------|------------------------|
| default   | node-server |             | http://127.0.0.1:49407 |
|-----------|-------------|-------------|------------------------|
🎉  Öffne Service default/node-server im Default-Browser...
❗  Weil Sie einen Docker Treiber auf darwin verwenden, muss das Terminal während des Ausführens offen bleiben.
```

- now check the response: `curl http://127.0.0.1:49407` should return `{"hello":"world"}%`

- stop minikube: `minikube stop`

### Notes

you could also try to run it locally using docker: `docker run -p 8080:8080 fhirter/node-server`

to delete deployment use: `kubectl delete deployment node-server`
