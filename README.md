# interview-assignment-sip
 
## How to Run? 

There are 2 entry point in the application

 ### SIP application (As per the pdf doc)

To run this application 

Simply navigate to the VS Code Run and Debug menu, and lunch 'Start SIP Application' from the preconfigured Launch.json

Or
Run the following from root folder of project
> cd ./internal/cmd/sip
> 
> go run ./main.go

this will initiate the server and once you see the message that server is listening you can send requests

### Client Simulation

This application is used to test the functionality of the SIP application with 5 concurrent clients, with 3 valid addresses and 2 invalid addresses

To run this application 

Simply navigate to the VS Code Run and Debug menu, and lunch 'Start Client Simulation' from the preconfigured Launch.json


Or
Run the following from root folder of project

> cd ./internal/cmd/client_simulation
> 
> go run ./main.go


## How to test?

This repo currently has unit test coverage for most of the functionalities, 

to test run the below command from the root of the project

> go test ./...


## Project structure

**internal** since this repo is meant only for internal features all code i sbundled under this folder

**pkg** this folder contains the code that can be shared among internal application

**data** conatins the raw data file provided as part of the assignemnet

**configs** contains app specific configs
**cmd** conatins the entry point for the applications (currently has 2)

**app** conatins the logic for the applications

**models** conatins the models for the application

**dal** data access layer , currently holds the in memory key value pair store using map data structure, can be replaced with other store without affecting the application

**services** contains the actual logic required for the application