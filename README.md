<h1 align="center">Algorithm API</h1>

<div align="center">

![STATUS](https://img.shields.io/badge/status-active-brightgreen?style=for-the-badge)
![LICENSE](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

</div>

---

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Testing](#testing)
- [Authors](#authors)

## 🔎 About <a name = "about"></a>

This project is meant to showcase the concepts of Backend Development and DevOps that I have learned so far. This project features API routing and design, CI/CD workflows and automations, unittesting, and containerization. Git & Github were used as the VCS. I dub this as "Algorithms-as-a-Service".

This program aims to create a Gin-based microservice that solves some classic programming problems. The Algorithm Microservice API provides a simple HTTP-based API to solve some of these problems. This microservice can be easily integrated into other applications or used for testing and learning purposes.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before running the Algorithm Microservice API, make sure you have the following prerequisites installed:

- Golang 1.18 or above

### Installing

Once you have ensured you have the appropriate version of Go installed, you can dive right in by cloning the repository.

```shell
git clone https://github.com/jgfranco17/go-with-the-flow.git
```

## 🚀 Usage <a name = "usage"></a>

### Dev mode

To run the API server in dev mode, simply execute either of the following commands:

```bash
# Default Go execution
go run service/cmd/main.go

# Justfile command runner
just run-debug <port>
```

### Calling the API

Once the service is up and running, we can now reach the endpoints. A basic example of using `curl` to reach endpoints is shown below. This example assumes that we are using the `v0` endpoints; see the documentation for more detailed guides on how the endpoint URL is determined.

```shell
$ curl 'localhost:8080/v0/algorithms/twosum/2-7-11-15/9'
{"Found":true,"Indices":[0,1]}
```

## 🔧 Testing <a name = "testing"></a>

### Running unittest suite

You can use the command runner to run a simple test suite execution using `go test`.

```shell
$ just test
```

## ✒️ Authors <a name = "authors"></a>

- [Chino Franco](https://github.com/jgfranco17)
