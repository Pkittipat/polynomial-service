package requests

type CalculateForm struct {
	X int `json:"x" binding:"required"`
	Y int `json:"y" binding:"required"`
	Z int `json:"z" binding:"required"`
}
