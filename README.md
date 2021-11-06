# go-minesweeper-api
API to play Mine Sweeper


Field:
```json
{
	"position": 1,
	"row": 1,
	"column": 0,
	"is_bomb": false,
	"bombs_around": 3
}
```

## Routes

POST -> `api/v1/start-game`

* Request Object:
```json
{
	"matrix": {
		"row": 7,
		"column": 4
	},
	"num_of_bombs": 6
}
```
POST -> `api/v1/play`

* Request Object:
```json
{
  "row": 2
  "column": 3
}
```

* Response Object:
```json
[<Fields>]
```
