package types

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
