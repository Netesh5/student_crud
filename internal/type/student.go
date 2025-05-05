package types

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
}
