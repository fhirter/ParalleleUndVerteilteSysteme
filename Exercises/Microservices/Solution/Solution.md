# Lösung Übung Microservices

Given the name of the cluster is `test` as set in `config.yaml`

## KIND

### Setup Cluster

- `brew install kind`
- create a cluster using [`config.yaml`](./config.yaml) with `kind create cluster --config=config.yaml`
- set context for kubectl: `kubectl cluster-info --context kind-test`
- List Clusters: `kind get clusters`

### Example Cluster

- Apply Ingress NGINX `kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/deploy-ingress-nginx.yaml`
- Apply ingress config `kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/usage.yaml`
- Test connection: `curl localhost/foo`,  `curl localhost/bar`

### Setup Service

- build docker image: `cd node && docker build -t node-server .`
- `kind load docker-image node-server -n test`
- list images running in kind: `docker exec -it test-control-plane crictl images`

### Notes

- kubectl apply does not seem to work on clusters

### minikube

[minikube.sigs.k8s.io/docs/start](https://minikube.sigs.k8s.io/docs/start/?arch=%2Fmacos%2Fx86-64%2Fstable%2Fbinary+download)

- `brew install minikube`
- `brew install kubectl`
- `minikube start`
- List all pods: `kubectl get pods -A`
- Install Minikube Dashboard: `minikube dashboard`