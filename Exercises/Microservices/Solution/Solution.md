# Lösung Übung Microservices

given the name of the cluster is `test` as set in `config.yaml`

## Setup Cluster

- `brew install kind`
- create `config.yaml` and create it with `kind create cluster --config=config.yaml`
- set context for kubectl: `kubectl cluster-info --context kind-test`
- List Clusters: `kind get clusters`

## Add Image

- build docker image: `(cd node && docker build -t node-server .)`
- `kind load docker-image node-server -n test`
- list images running in kind: `docker exec -it test-control-plane crictl images`
- Apply Ingress NGINX `kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/deploy-ingress-nginx.yaml`
- Apply ingress config `kubectl apply -f ingress.yaml`

## Delete Cluster

- `kind delete cluster --name kind-test`

## Offene Fragen

- Der Befehl `kubectl apply -f my-manifest-using-my-image:unique-tag` in https://kind.sigs.k8s.
  io/docs/user/quick-start/#loading-an-image-into-your-cluster ist unklar.

## Learnings

- kubectl apply does not seem to work on clusters created with kind
