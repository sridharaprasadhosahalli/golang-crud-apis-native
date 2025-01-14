# Building and Deploying a Simple CRUD API for items in a grocery store with Golang,Docker, and Kubernetes

This is a simple CRUD (Create, Read, Update, Delete) REST API for items in a grocery store with Go.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

- Go (at least version 1.23.4)

### Cloning the Project

```bash
git clone https://github.com/sridharaprasadhosahalli/golang-crud-apis-native.git
cd golang-crud-apis-native
```

### Initializing the Project

```bash
go mod init github.com/sridharaprasadhosahalli/golang-crud-apis-native
```

### Downloading Dependencies

```bash
go mod download
```

## Running the Application

```bash
go run main.go
```

The application will start running at `http://localhost:8080`.

## Run test cases 

```bash
go test
```


## API Endpoints

- **GET /items**: Fetch all items.
- **GET /items/:id**: Fetch a item by ID.
- **POST /items**: Create a new item. 
- **PUT /items/:id**: Update a item by ID.
- **DELETE /items/:id**: Delete a item by ID.

## Testing the API

You can use a tool like curl or Postman to test the API. Here are some examples:

- Fetch all items:

```bash
curl -X GET http://localhost:8080/items
```

Reponse 
```bash
[
  {
    "id": "1",
    "name": "Banana",
    "type": "Fruit",
    "price": 40
  },
  {
    "id": "2",
    "name": "Tomoto",
    "type": "Vegetable",
    "price": 30
  },
  {
    "id": "3",
    "name": "Cocumber",
    "type": "Vegetable",
    "price": 20
  }
]
```

- Fetch a item by ID:

```bash
curl -X GET http://localhost:8080/items/1
```
Response 

```bash
{
  "id": "1",
  "name": "Banana",
  "type": "Fruit",
  "price": 40
}
```
- Create a new item:

```bash
 curl http://localhost:8080/items \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","name": "potato","type": "vegetable","price": 49.99}'
```
Reponse
```bash
{
    "id": "4",
    "name": "potato",
    "type": "vegetable",
    "price": 49.99
}
```

- Update a item by ID:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": "3","name": "cabbage","type": "vegetable","price": 49.99}' http:/localhost:8080/items/3
```
Response
```bash 
{"id":"3","name":"cabbage","type":"vegetable","price":49.99}
```

- Delete a item by ID:
```bash
curl -X DELETE http://localhost:8080/items/3
```
Response
```bash
{"status":"deleted"}
```

## Dockerization of golang-crud-apis.
## Make sure docker desktop is installed .

## Build Docker image 

```bash
docker build --tag sridharaprasadhosahalli/golang-crud-apis-native .
```

## Check Built docker image 

```bash
docker image ls
```

## Run the Built docker image 

```bash
docker run -d -p 8080:8080 sridharaprasadhosahalli/golang-crud-apis-native
```

## verify docker container running

```bash
docker container ls
```

## Refer Testing the API section to test all the apis

## To run the test cases for Go in docker 

```bash
docker build -f Dockerfile -t golang-crud-apis-native-test --progress plain --no-cache --target run-test-stage .
```
## Install minikube for running kubernetes cluster and start the minikube preferably with docker as virtual machine

```bash
minikube start --driver=docker
```

## Install kubectl utility for cli

kubectl version

## creating deployment and service using kubectl 

```bash

kubectl apply -f deployment.yml

kubectl create -f service.yml

```
## check k8s resources came up and running

```bash
kubectl get po,svc,deployment,ep --show-labels
```

## List out the golang crud api service deployed to minikube to access it from browser by opening a minikube tunnel

```bash
minikube service --all
```

## Copy the url shown to access the golang crud apis in correspond to golang-crud-apis-svc and verify apis endpoint 
## example http://127.0.0.1:53356/items
```bash
minikube service --all 
|-----------|----------------------|-------------|---------------------------|
| NAMESPACE |         NAME         | TARGET PORT |            URL            |
|-----------|----------------------|-------------|---------------------------|
| default   | golang-crud-apis-svc |        8080 | http://192.168.49.2:31276 |
|-----------|----------------------|-------------|---------------------------|
|-----------|------------|-------------|--------------|
| NAMESPACE |    NAME    | TARGET PORT |     URL      |
|-----------|------------|-------------|--------------|
| default   | kubernetes |             | No node port |
|-----------|------------|-------------|--------------|
😿  service default/kubernetes has no node port
❗  Services [default/kubernetes] have type "ClusterIP" not meant to be exposed, however for local development minikube allows you to access this !
🏃  Starting tunnel for service golang-crud-apis-svc.
🏃  Starting tunnel for service kubernetes.
|-----------|----------------------|-------------|------------------------|
| NAMESPACE |         NAME         | TARGET PORT |          URL           |
|-----------|----------------------|-------------|------------------------|
| default   | golang-crud-apis-svc |             | http://127.0.0.1:53356 |
| default   | kubernetes           |             | http://127.0.0.1:53357 |
|-----------|----------------------|-------------|------------------------|
🎉  Opening service default/golang-crud-apis-svc in default browser...
🎉  Opening service default/kubernetes in default browser...
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.
```

