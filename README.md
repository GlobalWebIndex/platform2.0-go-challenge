# GlobalWebIndex Engineering Challenge - Elias Krontiris Submission

## Problem Analysis & Concerns

After reading carefully the challenge goals all assets (charts, insights, audience) could generate data from a single sample data-table. Guessing the Age, Gender, Country are `demographic information` and the only `metric` is the "hours spent on social media", I added the "Participants" schema and seeded mock data into this table.
Age(10-80), Gender(2) and Country(259) are randomly seeded values and "hours spent on social media" is based on a normal distribution with a spike on 0-2 hours daily, based on a personal estimation.

I should note that I fully understand that this is out of this challenge's scope and lacks performance, since all *Assets* are calculated based on live data and usually should be exported based on pre-calculated and transformed schemas.
Also noting that using a ORM deplets the neccesity to use repository pattern for storing data.

For convinience and easy testing the endpoint on every spin-up the DB is dropped and recreated with mock data.
I also include a collection of API calls test all chalenges endpoints [Postman Collection Link](https://www.getpostman.com/collections/dd8e929f0dd1124fbb3a)

## Docker Guidelines

To run the API use the following:

```bash
docker-compose -f ./docker-compose.yml
#In case Postgress delays 
docker-compose up -d --build
```

## API Routes table

Endpoint                                        | Description
------------                                    | -------------
(POST)/auth/signup                              | endpoint to create a user
(POST)/auth/login                               | authorize and return json web token
(GET)/assets                                    | based on the token returns all user's assets
(GET)/assets/{id}                               | get specific asset by id
(PATCH)/assets/{id}                             | edit asset's description or isFavorite
(POST)/assets/charts                            | create asset / chart
(POST)/assets/insigts                           | create asset / insight
(POST)/assets/audiences                         | create asset / audience

there are more endpoints that were made to make Seeding and Testing easier. All routes are specified on `/app/routes.go`

## Challenge Keypoints

- [x] A working server application with functional API is required
- [x] It is appreciated, though not required, if a Dockerfile is included.
- [x] Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time
- [x] Dockerfile included in the project
- [ ] Useful and passing tests would be also be viewed favourably *(Very Little Tests)

## Dependencies

- [x] Gorilla Mux - Router
- [x] GORM - A Go ORM
- [x] Dgrijalva JWT - JSON Web Token Library
- [x] Badoux Checkmail - Validate email on signup
- [x] Golang Crypto - Hashing passwords before storing  and Validity Comparison
- [x] Evanphx JSON Patch - Using a single endpoint to update asset title or asset favorite/unfavorite
