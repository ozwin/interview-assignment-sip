# interview-assignment-sip
## About
This project intended to demonstrate a simple SIP transaction lookup by its address, with a client-server model using a TCP socket
### Features 
+ Application can be initialized with a valid JSON transactions file (as long as it fits in the memory)
+ A client can make any number of requests on the same open connection as long as requests are terminated by a new line character(used as an EOL token , as per the doc)
+ A client is terminated by the server if there is no activity for more than 10 seconds
+ Configurable through internal/configs file
## Architecture
This project is using layered architecture, with each layer loosely coupled with other
<img width="500" alt="image" src="https://github.com/user-attachments/assets/e826498b-627e-445c-a083-ae4ab33a24c5">

**Data layer**: This layer is responsible for the data management and querying, for this project, I have defined an interface that ensures required functionalities are provided by the underlying data layer structs. This enables easy replacement of underlying data storage solutions, for this project although I am using simple in-memory Key Value pair implementation, it can be replaced with any persistent storage given that the solution implements the dal layer interfaces.

**Service layer**: This layer is responsible for the business logic of the application, it accesses the dal layer through the interface. This layer also ensures that there are no dependencies on the external user interfaces such as HTTP context, command line arg, buffer, etc and logic only depends on the function parameters, and the response is provided only through the returned values. 


**User Interface Layer/Presentation Layer**: This layer is responsible for accessing the application services. The same service layer can be used with any user interface.
 
## How to Run? 

There are 2 entry points in the repo

 ### SIP application (As per the pdf doc)

To run this application 

Navigate to the VS Code Run and Debug menu, and launch 'Start SIP Application' from the preconfigured Launch.json

Or
Run the following from the root folder of the project
> cd ./internal/cmd/sip
> 
> go run ./main.go

this will initiate the server and once you see the message that the server is listening you can send requests

### Client Simulation

This application is used to test the functionality of the SIP application with 5 concurrent clients, with 3 valid addresses and 2 invalid addresses

To run this application 

Navigate to the VS Code Run and Debug menu, and launch 'Start Client Simulation' from the preconfigured Launch.json


Or
Run the following from the root folder of the project

> cd ./internal/cmd/client_simulation
> 
> go run ./main.go

**p.s: Ensure you are running the SIP application first for this to work**

## How to test?

This repo currently has unit test coverage for most of the functionalities, 

to test run the below command from the root of the project

> go test ./...


## Project structure

**internal** Since this repo is meant only for internal features all code is bundled under this folder

**pkg** This folder contains the code that can be shared among internal application

**data** contains the raw data file provided as part of the assignment

**configs** contains app-specific configs
**cmd** contains the entry point for the applications (currently has 2)

**app** contains the logic for the applications

**models** contains the models for the application

**dal** data access layer, currently holds the in-memory key-value pair store using map data structure, can be replaced with other stores without affecting the application

**services** contains the actual logic required for the application
