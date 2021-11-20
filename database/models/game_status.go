package models

type GameStatus int

const (
	GameOver   GameStatus = iota - 1 // -1
	NotStarted                       // 0
	Playing                          // 1
	Winner                           // 2
)
