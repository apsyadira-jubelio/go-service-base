# go-service-base

This is an example of implementation of Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob. Read more at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

### What Are The Benefits ?

- [x] Easy to maintance
- [x] Easy to scalable you project
- [x] Readable code
- [x] Suitable for large projects or small projects
- [x] Easy to understand for junior or senior
- [x] And more

This project has 4 Domain layer :

- Domain Layer
- Repository Layer
- Usecase Layer
- Delivery Layer

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/apsyadira-jubelio/core-system.git

#move to project
$ cd core-system

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Stop
$ make stop

# Clean, build and run engine
$ make start
```

#### Run the Testing

```bash
$ make test
```

### Tools Used:

In this project, I use some tools listed below. But you can use any simmilar library that have the same purposes. But, well, different library will have different implementation type. Just be creative and use anything that you really need.

- ["github.com/vektra/mockery".](https://github.com/vektra/mockery) To Generate Mocks for testing needs.
