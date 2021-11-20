# go-minesweeper-api

This is an API created to study Go Lang.

The objective of it is to play Minesweeper. You can create a game through [this route](#start-the-game) and make a play through [this one](#make-a-play).

It's pretty simple, I hope you enjoy!!

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

### Start the Game
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

### Make a play
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


### Get some minesweep board already played or playing
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


## Errors

All the error response are `400 - BAD REQUEST`, following the format below:
```json
{
	"error": "Error String"
}
```