# Accounts-Service

Made a account microservice that can create user from a `POST` request and store it in a `Postgres DB`; and can get user information using `UUID` with a `GET` request. Used `bcrypt` to hash the password in the database.

NOTE - This isn't a Sign-Up and Sign-In service, it is just an account service although it can be turned into one with a few lines of code :)

## Running

1. Download the project in your GOPATH.
2. Create a database with name `users` or whatever name you like.
3. Go to `accounts/pkg/db/db.go` file and fill your database connection information to connect to the database.
4. cd inside the folder and run `docker-compose up` using the `docker-compose.yml` file or using `go run accounts/cmd/main.go`.
5. Start Postman for doing request and response as shown below or use curl commands.

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

## Database

 Useful commands to use -
 1. Switch over to the postgres account on your server by typing `sudo -i -u postgres`
 2. You can now access a Postgres prompt immediately by typing `psql`
 3. Use this to see all the tables `\dt`
 4. Use this to see info of a particular table `select * from "users";`
 
```
                  id                  |      email      |                           password                           
--------------------------------------+-----------------+--------------------------------------------------------------
 ce843823-e91f-4435-a4df-1b779ab2b587 | bcryptUser      | $2a$10$mtGwIQz3jhwINxH.0832hu5YQql/.kO.oJl3qJvo0EtwPqLZ.L0KO
 65c006f3-1c5c-47f1-bd72-48349ddcdb6a | email5     	| $2a$10$ah/zGl1JUvHZ9eUAAG4Nx.NWIdjI1yFVxQGAJ8Zpz4QKyL8U3ATH2
(2 rows)

```

5. To quit use `\q`
6. To exit use `exit`

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
|----------handler.go         Transport logic encode/decode data and  
|                             gorilla mux request reponse routing of the service.  
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
