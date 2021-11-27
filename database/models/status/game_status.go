package status

import "fmt"

type GameStatus int

const (
	GameOver   GameStatus = iota - 1 // -1
	NotStarted                       // 0
	Playing                          // 1
	Winner                           // 2
)

func (g GameStatus) ConvertToValue() string {
	value := fmt.Sprint(g)
	switch g {
	case GameOver:
		value = "Game Over"
	case NotStarted:
		value = "The game is ready to start"
	case Playing:
		value = "The game is currently being played"
	case Winner:
		value = "You Win!!"
	}
	return value
}
