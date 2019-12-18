# GlobalWebIndex Engineering Challenge

## The Challenge
Let's say that in GWI platform all of our users have access to a huge list of assets. We want our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following

Chart (that has a small title, axes titles and data)
Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month) e.g. Males from 24-35 that spent more than 3 hours on social media daily.
Build a web server which has some endpoint to receive a user id and return a list of all the user’s favourites. Also we want endpoints that would add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data. It’s up to you to decide the structure and we are not looking for something overly complex here (especially for the cases of audiences). There is no need to have/deploy/create an actual database although we would like to discuss about storage options and data representations.

Users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time.

## Installation
Clone the repository:
```
git clone https://github.com/arsotirchellis/platform2.0-go-challenge.git
cd platform2.0-go-challenge
```
If you have Docker installed just run the commands:

```sh
docker build -t gwi-demo .
docker run -p 8080:8080 gwi-demo
```

alternatively:
```sh
cd src
go build
go run main.go
```
>**go build** will download all the necessary dependencies declared on go module and will create an .exe file (on Windows)

### Help
You will find [here](https://github.com/arsotirchellis/platform2.0-go-challenge/blob/master/GWI-Challenge-Postman-Collection.postman_collection.json) a postman collection JSON file which you can import into a Postman client.

>If you have an unauthorized response code, you have to make a login call and copy, paste the JWT access token in the authorization header as BEARER your_token.

### Todos
* Dependency Injection for every service, repository, controller.. [Wire](https://github.com/google/wire) is a good solution. Also check [this article](https://blog.drewolson.org/dependency-injection-in-go).
* After the implementation of the pattern above, you have to make 80% unit test coverage at least on service layer.
* Return asset's user description on asset dtos
* Encrypt the asset id, and decrypt it on every request. Don't expose database primary key on clients. Title is not recommended also to be on url as param.

Happy reviewing!
