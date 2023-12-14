<div align="center">
<h1>Echo Scaffold</h1>
<img src="https://skills.thijs.gg/icons?i=golang,docker,postgresql,linux,vscode,github,markdown" />
</div>

## Description

This project is a simple scaffold template for an Echo application.

## Authors

By order of contribution:

- Gabriel Santos Cardoso - [eng.gabrielscardoso@gmail.com](mailto:eng.gabrielscardoso@gmail.com)

## Installation

1. First of all, make sure to have the Golang SDK installed in your local environment, as well as Docker and Docker Composer
2. Clone the repository in your workspace and install the dependencies using the command `go mod tidy`
3. Copy the `.env.example` as `.env` and up the services using `docker-compose up -d` (the support for other databases still under development)
4. Optionally, you can change the name of application from `echo-scaffold` for the name of your application
5. After all steps above your application is already to be launched with `go run cmd/main/main.go` 🌟

## Licence

This project is licenced under the [MIT Licence](LICENSE)
