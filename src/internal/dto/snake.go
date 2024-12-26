package dto

type Snake struct {
	Start int `json:"start"` // Start position (snake's head)
	End   int `json:"end"`   // End position (snake's tail)
}
