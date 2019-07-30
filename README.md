# go-api  - Go lang webserver assignment


This project was created to demonstrate skills in GO lang


## Running locally

git clone git@github.com:tsolman/go-api.git

cd go-api
go build && ./go-api

```
i.e. 
open POSTMAN and GET http://localhost:8000/api/health
```

## Available endpoints

GET "/api/health"  - Chech app health

POST "/api/user/" - Create user
##### Sample input
```
{
	"name": "Ioannou",
	"asset": [{
        "description":"new insight",
        "assetype":"insight",
        "isfav": false,
        "data": {
            "text": "30% of the univerce is mine"
        }
    }]
}
```

GET "/api/user/{id}" - Get user by id

PUT "/api/asset/{userid}" - Create asset per user

##### Sample input
```
{
	"description":"new insight",
	"assetype":"insight",
	"isfav": false,
	"data": {
		"text": "30% of the univerce is mine"
	}
}
```
GET "/api/assets" - Get assets

PUT "/api/user/{userid}/asset/{id}" - Update asset description
##### Sample input
```
{
	"description" : "new description"
}
```

PUT "/api/user/{userid}/asset/{id}/fav - Mark asset as favourite
##### Sample input
```
{
	"isfav" : false
}
```


## Deployment

The app is live at https://golang-api-test.herokuapp.com/

## Built With

* [GOLANG](https://golang.org/) - The GO programming language

* [MUX](https://www.gorillatoolkit.org/pkg/mux) - Gorilla web toolkit


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
