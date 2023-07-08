# Albo

Syncs marvel comics to local and we expose the data.

## Technology

[Colaboratos service](/services/colaborator/README.md)
- Golang (Gin)
- NoSQL (mongodb)

[Characters service](/services/character/README.md)
- Node (Nestjs)
- NoSQL (mongodb)

## Environment variables 

| Variable  | Type  | Description  | Default  |
|---|---|---|---|
| MONGO_URI | string | mongo uri string  | mongodb://mongo_albo:27017 |
| DURATION |  int | interval time to sync (minutes) | 60 |
| DB_NAME | string | db name | albo |

`No needed to set something at the moment`


## Requirements

* docker
* docker-compose

## Run project

```sh
# Clone repo
$ git clone git@github.com:BraulioAguilarDev/albo.git

# Go to albo directory
$ cd albo

# Create and run the containers
$ make dc-up

# Remove all
$ make dc-down
```

## Test endpoints

Note: For sharing the same endpoint `localhost` we are using a nginx proxy_pass

`GET: http://localhost/marvel/colaborators/ironman`

```json
{
	"last_sync": "08-07-2023 18:21:51",
	"editors": [
		"Tom Brevoort"
	],
	"writers": [
		"Peter Milligan",
		"Michael Avon Oeming",
		"J. Michael Straczynski",
		"John Miller",
		"Mark Ricketts",
		"Chuck Austen",
		"Stan Lee",
		"Brian Michael Bendis"
	],
	"colorists": [
		"Liquid! Color",
		"Jason Keith",
		"Wil Quintana",
		"Laura Villari",
		"Paul Mounts",
		"J. D. Mettler",
		"Edgar Delgado",
		"Christopher Sotomayor",
		"Frank D'ARMATA"
	]
}
```

`GET: http://localhost/marvel/characters/ironman`

```json
{
	"last_sync": "Sat Jul 08 2023 18:21:53 GMT+0000 (Coordinated Universal Time)",
	"characters": [
		{
			"character": "Wolverine",
			"comics": [
				"Wolverine Saga (2009) #7",
				"Secret Invasion Infinity Comic (2023) #18"
			]
		},
		{
			"character": "Avengers",
			"comics": [
				"Avengers Unlimited Infinity Comic (2022) #53",
				"Avengers Unlimited Infinity Comic (2022) #52",
				"Avengers Unlimited Infinity Comic (2022) #51",
				"Secret Invasion Infinity Comic (2023) #15",
				"Secret Invasion Infinity Comic (2023) #16",
				"Secret Invasion Infinity Comic (2023) #18",
				"Avengers Unlimited Infinity Comic (2022) #50",
				"Avengers Unlimited Infinity Comic (2022) #49",
				"Avengers: Electric Rain Infinity Comic (2022) #13",
				"Avengers Unlimited Infinity Comic (2022) #30",
				"Avengers: Electric Rain Infinity Comic (2022) #12",
				"Avengers: Electric Rain Infinity Comic (2022) #11"
			]
		},
		{
			"character": "Captain America",
			"comics": [
				"Avengers Unlimited Infinity Comic (2022) #53",
				"Avengers Unlimited Infinity Comic (2022) #52",
				"Secret Invasion Infinity Comic (2023) #15",
				"Secret Invasion Infinity Comic (2023) #18"
			]
		},
		{
			"character": "Serpent Society",
			"comics": [
				"Avengers Unlimited Infinity Comic (2022) #53"
			]
		}
	]
}
```

## TODO

- Testing: Like we are using interfaces we can implement testing with [testify](https://github.com/stretchr/testify).

- Documentation: I would propose using [swagger](https://swagger.io/).
