# Travelex Pre Interview Application #
## About ##
This repository contains a simple microservice application written in [go](http://golang.com).

## Requirements ##
You'll need docker and a git client to run this application.

## Running ##
```bash
git clone https://github.com/pennywisdom-other/travelex.git
cd travelex
docker build -t travelex-test .
docker run -it -p 8080:8080 travelex-test
```

This will take a minute or 2 (depending on your connection speed) and will compile the application then run the microservice as a docker container.
The documentation will be logged in the output.
You can then use something like [postman](https://www.getpostman.com) to interact with it.

## API ##

### V1 ###
The API for this service supports a single endpoint that can be queried using a 2 querystring options.

#### Endpoint ####

/v1/countries?target=source
/v1/countries?target=destination

You can query the service as follows:

```bash
$ curl -- header "Accept:application/json" http://localhost:8080/v1/countries?target=source
```

```bash
$ curl -- header "Accept:application/json" http://localhost:8080/v1/countries?target=destination
```

#### Request Headers ####
The API supports requests made with the following request headers:

```
Accept:application/json
Content-Type:application/json
```

Unsupport values will receive a 400 Bad Request reponse.
