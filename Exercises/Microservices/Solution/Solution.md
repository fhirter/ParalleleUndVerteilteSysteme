# Lösung Übung Microservices

given the name of the cluster is `test` as set in `config.yaml`

- `brew install kind`
- create `config.yaml` and create it with `kind create cluster --config=config.yaml`
- set context for kubectl: `kubectl cluster-info --context kind-test`
- List Clusters: `kind get clusters`
- build docker image: `cd node && docker build -t node-server .`
- `kind load docker-image node-server -n test`
- list images running in kind: `docker exec -it test-control-plane crictl images`
- kubectl apply does not seem to work on clusters
- Apply Ingress NGINX `kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/deploy-ingress-nginx.yaml`
- Apply ingress config `kubectl apply -f ingress.yaml`