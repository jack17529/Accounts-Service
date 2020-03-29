# Accounts-Service

## Running

1. Download the project in your GOPATH.
2. Go to `accounts/pkg/db/db.go` file and fill your database connection information to connect to the database.
3. cd inside the folder and run `docker-compose up` using the `docker-compose.yml` file or using `go run accounts/cmd/main.go`.
4. Start Postman for doing request and response as shown below or use curl commands.

## Request And Response

1. POST request done on `http://localhost:8081/create-user`

```
{
	"User":{
		"email":"email5", 
		"password":"password5"
	}
}
```

POST response

```
{
    "Id": "65c006f3-1c5c-47f1-bd72-48349ddcdb6a",
    "Error": null
}
```

2. GET request done on `http://localhost:8081/get-user`

```
{
	"ID" : "65c006f3-1c5c-47f1-bd72-48349ddcdb6a"
}
```

GET response

```
{
    "Email": "email5",
    "Error": null
}
```

## Architechture
```
accounts/  
|---cmd/  
|------service/  
|----------server.go          Wire the service.  
|----------server_gen.go      Also wire the service.  
|------main.go                Runs the service  
|---pkg/  
|------endpoints/  
|----------endpoint.go        The endpoint logic along with structures for request and reponse.  
|----------endpoint_gen.go    This will wire the endpoints.  
|----------middleware.go      Endpoint middleware  
|------http/  
|----------handler.go         Transport logic encode/decode data and gorilla mux request reponse routing of the service.  
|----------handler_gen.go     This will wire the transport.  
|------io/  
|----------io.go              The input output structs.  
|------db/  
|----------db.go              Gets connection to PostgreSQl database.  
|------service/  
|----------middleware.go      The service middleware.  
|----------service.go         Business logic.
```

## Advantages of UUID

1. Can generate them offline.
2. Makes replication trivial (as opposed to int's, which makes it REALLY hard)
3. ORM's usually like them
4. Unique across applications. So We can use the PK's from our CMS (guid) in our app (also guid) and know we are NEVER going to get a clash.
