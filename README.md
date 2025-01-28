# web-crawler-go

Web-Crawler-Go is an open-source web crawling software designed to extract data from websites by mentioning the url and depth in the main.go file before running the application.

# Prerequisites

- Go 1.23.3+

Installation Guide
To install Task Manager Golang, follow these steps:

Step 1: Install Go
Task Manager Golang is built using Go, so you need to have Go installed on your system. You can download the latest version of Go from the official Go website (https://golang.org/dl/). Follow the installation instructions provided on the website to install Go on your system.

Step 2: Install Go Modules
Task Manager Golang uses Go Modules to manage dependencies. To use Go Modules, you need to have Go 1.11 or later version installed on your system. If you haven't installed Go 1.11 or later, install Go using the instructions provided in step 1 and then follow the instructions below.

Step 3: Clone the Repository
Clone the Task Manager Golang repository from GitHub using the following command:

```shell
https://github.com/mohanrajnambe/web-crawler-go.git
cd web-crawler-go
```

Step 4: Install all the dependencies

```shell
go mod tidy
```

Step 5: Build and Install

```shell
go build -o web-crawler-go
```

Step 6: Run Web crawler

```shell
./web-crawler-go
```
