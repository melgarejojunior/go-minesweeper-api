package entities

type CreateMinesweeper struct {
	Matrix     `json:"matrix"`
	NumOfBombs int `json:"num_of_bombs"`
}

type Matrix struct {
	Row    int `json: "row" binding:"required`
	Column int `json: "column" binding:"required`
}
