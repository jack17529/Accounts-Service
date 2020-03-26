# Accounts-Service

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

2. GET request done on `http://localhost:8081/create-user`

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
