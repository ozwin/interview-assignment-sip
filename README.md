# interview-assignment-sip
 
## How to Run? 

There are 2 entry point in the application

 ### SIP application (As per the pdf doc)

To run this application 

Simply navigate to the VS Code Run and Debug menu, and lunch 'Start SIP Application' from the preconfigured Launch.json

Or

navigate to the folder internal/cmd/sip and run 

go run ./main.go

this will initiate the server and once you see the message that server is listening you can send requests

### Client Simulation

This application is used to test the functionality of the SIP application with 5 concurrent clients, with 3 valid addresses and 2 invalid addresses

To run this application 

Simply navigate to the VS Code Run and Debug menu, and lunch 'Start Client Simulation' from the preconfigured Launch.json


Or

navigate to the folder internal/cmd/client_simulation and run 

go run ./main.go


## How to test?

This repo currently has unit test coverage for most of the functionalities, 

to test run the below command from the root of the project

go test ./...

