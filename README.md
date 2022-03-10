# Multi-tenancy on Kubernetes
Automatic deployment of tenant applications to k8s separated by namespaces

## Structure

### App
Tenant application (Go REST API)

### Helm Chart
K8s template for tenant application

- Rest API
- MySQL DB

### Provisioner
REST API responsible for deploying tenant application into a specified namespace