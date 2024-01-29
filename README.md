# GeoIP Service

This project is an example implementation of a GeoIP service using gRPC and gRPC-Gateway in the Go programming language.


## Installation

1. Clone the repository:
```bash
git clone https://github.com/woozie-10/geoip-service.git
```

## Running the Server

1. Make sure you have the GeoIP2-City.mmdb file, which contains the GeoIP database. This file should be placed in the root directory of the project.
2. Start the server:
```bash
	docker-compose up --build
```
## Sending Requests
You can send requests to the server using gRPC or HTTP.
1. To send a gRPC request, use client code in Go language or any other gRPC-supporting library.
2. To send an HTTP request, use the **curl** utility or other tools for sending HTTP requests:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"address":"127.0.0.1"}' http://localhost:5052/v1/ip
```
## Contribution
If you have any suggestions or fixes, please create an issue or pull request. I welcome your contributions!
