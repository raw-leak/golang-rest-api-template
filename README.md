
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

## Features

- List key features and functionalities of your API.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [Solution notes](#solution-notes)
- [License](#license)

## Installation

Provide instructions on how to install your Golang REST API. Include any prerequisites and steps to set up the environment.

```bash
# Example installation steps
$ git clone https://github.com/yourusername/your-api.git
$ cd your-api
$ go build
$ ./your-api
```

## Usage

Explain how to use your API. Include code examples, endpoints, and any required parameters.

```bash
# Example usage
$ curl -X GET http://localhost:8080/api/endpoint
```

## API Documentation

Provide a link to detailed API documentation, if available.

## Configuration

Explain how to configure your API, including environment variables, configuration files, or any other relevant settings.

```bash
# Example configuration
$ export API_KEY=your-api-key
$ ./your-api
```

## Contributing

Tell others how they can contribute to your project. Include guidelines for submitting issues or pull requests.

## Solution notes

- :trident: clean architecture (handler->service->repository)
- :book: standard Go project layout
- :cd: github CI/CD + docker compose + Makefile included
- :card_file_box: MongoDB connection included
- :white_check_mark: tests with mocks included

## License
This project is licensed under the [MIT License](LICENSE).

