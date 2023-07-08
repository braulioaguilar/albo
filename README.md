# Albo

Syncs marvel comics to local and we expose the data.

## Technology

[Colaboratos service](/services/colaborator/README.md)
- Golang (Gin)
- NoSQL (mongodb)

[Characters service]((/services/characters/README.md))
- Node (Nestjs)
- NoSQL (mongodb)

## Environment variables 

| Variable  | Type  | Description  | Default  |
|---|---|---|---|
| MONGO_URI | string | mongo uri string  | mongodb://mongo_albo:27017 |
| DURATION |  int | interval time to sync (minutes) | 60 |
| DB_NAME | string | db name | albo |


## Requirements

* docker
* docker-compose

## Run project

```sh
# Clone repo
$ git clone git@github.com:BraulioAguilarDev/albo.git

# Go to albo directory
$ cd albo

# Copy the .env file then replace all ## change me ## 
$ make env

# Create and run the containers
$ make dc-up

# Remove all
$ make dc-down
```
