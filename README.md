# go-minesweeper-api
API to play Mine Sweeper


## Entities

* **Field**
```json
{
	"row": 1,
	"column": 0,
	"is_bomb": false,
	"bombs_around": 3
}
```

* **Minesweeper**
```json
{
    "id": 20,
    "row": 3,
    "column": 3,
    "num_of_bombs": 3,
    "created_at": "2021-11-20T16:42:10.45592-04:00"
}
```

* **Game**
```json
{
    "id": 13,
    "fields": [<Field>],
    "minesweeper": <Minesweeper>,
    "game_status": -1
}
```

* **GameStatus**
	GameOver   		-1
	NotStarted		0
	Playing    		1
	Winner			2


## Routes

**POST** -> `api/v1/minesweeper/start`

* Request Object:
```json
{
	"matrix": {
		"row": 3,
		"column": 3 
	},
	"num_of_bombs": 3
}
```

* Response Example
```json
{
    "id": 15,
    "fields": [
        {
            "row": 0,
            "column": 0,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 0,
            "column": 1,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 0,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 1,
            "column": 0,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 1,
            "column": 1,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 1,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 2,
            "column": 0,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 2,
            "column": 1,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 2,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": false
        }
    ],
    "minesweeper": {
        "id": 22,
        "row": 3,
        "column": 3,
        "num_of_bombs": 3,
        "created_at": "2021-11-20T17:10:05.915472-04:00"
    },
    "game_status": 0
}
```

**POST** -> `api/v1/game/{game_id}/play`

* Request Object:
```json
{
  "row": 2,
  "column": 0
}
```

* Response Example
```json
{
    "id": 15,
    "fields": [
        {
            "row": 0,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": true
        },
        {
            "row": 2,
            "column": 0,
            "is_bomb": false,
            "bombs_around": 2,
            "is_opened": true
        }
    ],
    "minesweeper": {
        "id": 22,
        "row": 3,
        "column": 3,
        "num_of_bombs": 3,
        "created_at": "2021-11-20T17:10:05.915472-04:00"
    },
    "game_status": 1
}
```

**GET** -> `api/v1/game/{game_id}`

If the game still playing:

* Response Example
```json
{
    "id": <game_id>,
    "fields": [
        {
            "row": 0,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 0,
            "is_opened": true
        },
        {
            "row": 2,
            "column": 0,
            "is_bomb": false,
            "bombs_around": 2,
            "is_opened": true
        }
    ],
    "minesweeper": {
        "id": 22,
        "row": 3,
        "column": 3,
        "num_of_bombs": 3,
        "created_at": "2021-11-20T17:10:05.915472-04:00"
    },
    "game_status": 1
}
```

If the game has already ended:
* Response Example
```json
{
    "id": <game_id>,
    "fields": [
        {
            "row": 0,
            "column": 0,
            "is_bomb": true,
            "bombs_around": 1,
            "is_opened": true
        },
        {
            "row": 0,
            "column": 1,
            "is_bomb": true,
            "bombs_around": 0,
            "is_opened": false
        },
        {
            "row": 0,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 2,
            "is_opened": false
        },
        {
            "row": 1,
            "column": 0,
            "is_bomb": false,
            "bombs_around": 3,
            "is_opened": false
        },
        {
            "row": 1,
            "column": 1,
            "is_bomb": true,
            "bombs_around": 2,
            "is_opened": false
        },
        {
            "row": 1,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 2,
            "is_opened": false
        },
        {
            "row": 2,
            "column": 0,
            "is_bomb": false,
            "bombs_around": 1,
            "is_opened": false
        },
        {
            "row": 2,
            "column": 1,
            "is_bomb": false,
            "bombs_around": 1,
            "is_opened": false
        },
        {
            "row": 2,
            "column": 2,
            "is_bomb": false,
            "bombs_around": 1,
            "is_opened": true
        }
    ],
    "minesweeper": {
        "id": 20,
        "row": 3,
        "column": 3,
        "num_of_bombs": 3,
        "created_at": "2021-11-20T16:42:10.45592-04:00"
    },
    "game_status": -1
}
```


## Errrors

All the error response are `400 - BAD REQUEST`, following the format below:
```json
{
	"error": "Error String"
}
```