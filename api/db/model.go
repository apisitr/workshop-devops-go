package db

type Account struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ServedBy string `json:"servedBy"`
}

type AccountData struct {
	ID   string `json:""`
	Name string `json:"name"`
}