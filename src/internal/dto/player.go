package dto

type Player struct {
	Name       string `json:"name"`       // Player's name
	CurrentPos int    `json:"currentPos"` // Current position of the player on the board
}
