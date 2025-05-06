package types

type Student struct {
	Id      int    `json:"id" `
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Age     int    `json:"age" validate:"required"`
	Email   string `json:"email" validate:"required"`
}
