## Starting the Web Server

Using docker the web server can start at port 8080 as: 

```bash
docker compose up
```
## Executing the Test Cases

Make sure the web server is running on background and execute the tests as:

```bash
docker-compose up -d
docker exec platform20-go-challenge_app_1 /bin/bash -c  "export CGO_ENABLED=0 && go test -v"
```

## API Points
A basic authentication is used, so the necessary username and password must also be provided in the request.

* username: admin
* password: 12345

The admin credentials are only applicable in the admin related requests, for any user related requests the username and password of the respective user should be used. `Hint: using the admin requests one can obtain everything`

#### Import Postman Collection

A postman collection with all the api requests available: **golang_assignment.postman_collection.json**

## The following RESTful endpoints are provided:

####  Health check endpoint, (without authentication)
* `GET /ping`: a ping service

####  Admin authenticated endpoints
* `POST /assets/add/{num}`: generate random assets. 
* `GET /assets`: returns all assets 
* `GET /users`: returns all assets 
* `GET /asset/{assetid}`: returns a specific asset based on its id.

#### User authenticated endpoints
* `GET /user/{userid}`: all assets of user
* `DELETE /user/{userid}/asset/{assetid}`: removal of specific asset of user
* `PUT /user/{userid}/asset/{assetid}/favor`: mark asset as favorite
* `PUT /user/{userid}/asset/{assetid}/unfavor`: unmark asset as favorite
* `PUT /user/{userid}/asset/{assetid}/editdesc`: edit asset's description
