# Multi-tenancy on Kubernetes
Automatic deployment of tenant applications to k8s separated by namespaces

## Structure

### App
Tenant application
- REST API
- MySQL DB

### Helm Chart
K8s template for tenant application


### Provisioner
REST API responsible for deploying tenant application into a specified namespace

## Requirements

- Helm

- Minikube cluster running locally

- Kubectl

## Setup

1. Run Provisioner

2. `cd provisioner`

3. `go run main.go`

Provisioner runs on `localhost:7000`

## Creating tenants using Provisioner API

### /POST

request body:

ns - Namespace to deploy tenant application

port - Nodeport for service  (30000-32768)


```json
{
    ns: "namespace"
    port: "30000"
}
```

### View created namespaces

`kubectl get namespace`

### View pods in created namespace

`kubectl get pods -n [namespace]`

### View services  in created namespace

`kubectl get services -n [namespace]`

### Provide external IP to deployment

`minikube service clients -n [namespace]`
