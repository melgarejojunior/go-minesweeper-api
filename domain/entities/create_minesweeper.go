package entities

type CreateMinesweeper struct {
	matrix     `json:"matrix"`
	NumOfBombs int `json:"num_of_bombs"`
}

type matrix struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}
