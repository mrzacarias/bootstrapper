# Bootstrapper

Boilerplate/Scaffolding tool to bootstrap new k8s applications

## Setup

1. Install Docker (https://www.docker.com/)
1. `./script/setup.sh` will update/download/generate needed resources and generate Docker container
1. `./script/test.sh (test_folder)` will run tests
1. `./script/run.sh` will compile the binary and run

## The Boilerplate Application

Bootstrapper's boilerplate application was created trying to cover the most common operations a k8s application will need to handle.

There are currently 2 types of services to chose:
1. Stateless HTTP Endpoint (default)
1. Stateful HTTP Endpoint (With Postgres scaffolding)

### Stateless HTTP Endpoint

Basic HTTP server with:
1. **Root**: Will just present the service
1. **Health Check Endpoints**: Used by Kubernetes to check the service liveness
1. **Emoji**: Fun endpoint that will receive a GET request with a `name` param and return the correspondent GitHub emoji (if existent)

#### Creating app
```
$ bootstrapper --app_name=stateless --app_prefix=STL --app_template=1
```

#### Example
```
http://localhost:1234/emoji?name=shipit
```

### Stateful HTTP Endpoint

Same as the Stateless version, but with basic Postgres configuration for:
1. Setup postgres database images (docker-compose)
1. Run the migrations during the image creation (initdb.d)
1. Packages and tests updated to record emoji requests
1. Last 5 requests on emoji feedback page

#### Creating app
```
$ bootstrapper --app_name=stateful --app_prefix=STF --app_template=2
```

## K8s configuration

### deployment.yml

1. The number of replicas (`1`)
1. Image for the main container on the pod (`TBD by user`)
1. Main container resourcing (CPU and Memory) (`100m and 100Mi`)
1. Environment variables

### horizontal_autoscaler.yml

1. `minReplicas` and `maxReplicas` (`1` and `4`)
1. Target values to increase or reduce the number of replicas (`cpu avg 50%` and `memory avg 50%`)

### ingress.yml

1. Ingress configuration for HTTPS host (`TBD by user`)

### network_policy.yml

1. NetworkPolicy for ports `8080` and `8081`

### pod_disruption_budget.yml

1. PodDisruptionBudget with default rule `maxUnavailable: 25%`

### service.yml

1. ClusterIP service for metrics (Service port `81` mapped to Container port `8081`)
1. NodePort service for open endpoints (e.g. emoji) (Service port `80` mapped to Container port `8080`)
