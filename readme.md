# Technical Exercise

A simple microservice in Go, exposed through gRPC, applying Docker and Kubernetes. Still in development.

The twelve-factor app best practices and how they fit with this project:
1.- Codebase:
Self-descriptive, this service has been developed using git as the only version-control system.
2.- Explicitly declare and isolate dependencies:
I declare all the dependencies of the package and resolve them using go modules.
3.- Store configuration in the environment:
The only thing that we've got to configure is the address of the service, both in the server and the client. This 
