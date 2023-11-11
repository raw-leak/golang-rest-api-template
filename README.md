
<div align="center">
 <h1> GOLANG REPO TEMPLATE </h1>
</div>

<div align="center">

  [![Go](https://img.shields.io/badge/Go-v1.21-blue.svg)](https://golang.org/)
  ![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/MykhayloGusak/golang-rest-api-template/main)
  ![GitHub Repo stars](https://img.shields.io/github/stars/MykhayloGusak/golang-rest-api-template)
  ![License](https://img.shields.io/badge/license-MIT-green)
  [![CI](https://github.com/MykhayloGusak/golang-rest-api-template/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/MykhayloGusak/golang-rest-api-template/actions/workflows/ci.yaml)
  ![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/MykhayloGusak/golang-rest-api-template/main)
  ![Github Repository Size](https://img.shields.io/github/repo-size/MykhayloGusak/golang-rest-api-template)
  ![GitHub contributors](https://img.shields.io/github/contributors/MykhayloGusak/golang-rest-api-template)
  ![Github Open Issues](https://img.shields.io/github/issues/MykhayloGusak/golang-rest-api-template)
</div>


<div align="center">
  <img src="images/go-logo.png" alt="Project Logo" width="200">
</div>

## Introduction

Briefly introduce your Golang REST API project here. Provide context about what problem it solves and why it's valuable.

## Table of Contents

- [Features](#features)
- [Usage](#usage)
- [Configuration](#configuration)
- [License](#license)

## Features

- 🔱 Clean architecture (handler->service->repository)
- 📖 Standard Go project layout
- 💿 CI/CD with GitHub Actions
- 🐳 Docker compose 
- ⚙️ Makefile 
- 🗃️ MongoDB connection
- ✔️ Unit/E2E tests with mocks and testify
- 🚦 Graceful shutdown

## Prerequisites

Before you can run this project, make sure you have the following prerequisites installed:

1. **Go 1.21+**: You'll need Go 1.21 or a later version installed. You can download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)

2. **MongoDB**: You'll need a MongoDB server running locally or accessible from your development environment. You can download and install MongoDB from the official website: [https://www.mongodb.com/try/download/community](https://www.mongodb.com/try/download/community)

3. **Make**: Make sure you have Makefile installed. Most Unix-like systems have it pre-installed. For Windows, you can use tools like [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm).

OR

- **Docker**: If you prefer a containerized environment that includes MongoDB, you can use Docker. You can download and install Docker from the official website: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)

   Additionally, this project includes a Docker Compose file that sets up the necessary services, including MongoDB, to run the application. Make sure you have Docker Compose installed as well: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)

Once you have these prerequisites in place, you can proceed with setting up and running the project.

## Usage

TODO

Explain how to use your API. Include code examples, endpoints, and any required parameters.

```bash
# Example usage
$ curl -X GET http://localhost:8080/api/endpoint
```

## Configuration

TODO

Explain how to configure your API, including environment variables, configuration files, or any other relevant settings.

```bash
# Example configuration
$ export API_KEY=your-api-key
$ ./your-api
```

## License
This project is licensed under the [MIT License](LICENSE).

