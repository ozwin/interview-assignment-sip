# interview-assignment-sip
 
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
