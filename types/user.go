package types

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
type Member struct {
	Name   string
	Age    int
	Active bool
}
