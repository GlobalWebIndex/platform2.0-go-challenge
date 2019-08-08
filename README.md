# GlobalWebIndex Engineering Challenge

## Introduction
This repo contains the source for Skyritsis GWI  Engineering Challenge.<br>

## Scope
Build a web server which has some endpoint to receive a user id and return a list of all the user’s assets. Also we want endpoints that would add an asset to favourite, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data. It’s up to you to decide the structure and we are not looking for something overly complex here (especially for the cases of audiences). There is no need to have/deploy/create an actual database although we would like to discuss about storage options and data representations.

## Architecture
The solution provided is loosely based on Uncle Bob’s Architecture (Golang Clean Architecture).

## Installation

In order to intall and run this app, you need to first get the source code from github.<br>After that, there are two ways to run in. Using docker or golang directly.
### Docker Method
```
$ cd /path/to/gwiChallenge
$ docker build --tag=gwi .
$ docker run -p 1323:1323 --name gwi gwi
#optional api tests
$ docker run --net=host -it gwi /app/test
```

### Golang Method
Requires Go 1.12 and go modules
```
$ cd /path/to/gwiChallenge
$ go build && ./gwi 
#optional api tests
$ cd /path/to/gwiChallenge/apiTests
$ go build && ./apiTests 
```

## API Description

All API calls have JSON request and response bodies
<br>
- - - -
### Get All System Users

```http
GET host_ip:1323/api/users
```
#### Response

```javascript
[
    {
        "id": 1
    },
    .
    .
    .
]
```
- - - -
### Create System User

```http
GET host_ip:1323/api/users/create
```

#### Response created user

```javascript
{
    "id": 1
}
```
- - - -
### Get a System User

```http
GET host_ip:1323/api/users/get/:id
```
#### Response retrieved user

```javascript
{
    "id": 1
}
```
- - - -
### Delete a System User

```http
DELETE host_ip:1323/api/users
```

#### Request
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `int` | **id of the user to be deleted** |

#### Response deleted user

```javascript
{
    "id": 1
}
```
- - - -
### Get User Favourites

```http
POST host_ip:1323/api/users/favourites
```

#### Request
```javascript
{
    "user":{
        "id":1
    },
    "next_token": 0,
    "page_size": 100
}
```

#### Response UserFavourites Paginated

```javascript
{
    "user": {
        "id": 1,
        "assets": {
            "550": {
                "id": 550,
                "description": "XVlBzgbaiC",
                "insight": {
                    "insight": "MRAjW"
                }
            },
            .
            .
            .
        }
    }
}
```
- - - -
### Add User Favourite

```http
PUT host_ip:1323/api/users/favourites
```

#### Request
```javascript
{
    "user":{
        "id":1
    },
    "asset":{
        "id":1
    }
}
```

#### Response Added Favorite ID

```javascript
{
    "id": 1
}
```
- - - -
### Remove User Favourite

```http
DELETE host_ip:1323/api/users/favourites
```

#### Request
```javascript
{
    "user":{
        "id":1
    },
    "asset":{
        "id":1
    }
}
```

#### Response User ID

```javascript
{
    "id": 1
}
```
- - - -
### Get An Asset

```http
GET host_ip:1323/api/assets/:id
```

#### Response Asset

```javascript
{
    "id": 1,
    "description": "...",
    "chart": {
        "title": "...",
        "axis_titles": [
            "...",
            "..."
        ],
        "data": [
            [
                "...",
                "...",
            ]
        ]
    }
}
```
- - - -
### Create An Asset

```http
POST host_ip:1323/api/assets/
```

#### Request Asset

```javascript
{
    "description": "...",
    "chart": {
        "title": "...",
        "axis_titles": [
            "...",
            "..."
        ],
        "data": [
            [
                "...",
                "...",
            ]
        ]
    }
}
```
#### Response Asset

```javascript
{
    "id": 1,
    "description": "...",
    "chart": {
        "title": "...",
        "axis_titles": [
            "...",
            "..."
        ],
        "data": [
            [
                "...",
                "...",
            ]
        ]
    }
}
```
- - - -
### Update An Asset Description

```http
POST host_ip:1323/api/assets/
```

#### Request Asset

```javascript
{
    "id": 1,
    "description": "...",
}
```
#### Response Asset

```javascript
{
    "id": 1,
    "description": "..."
}
```
- - - -
