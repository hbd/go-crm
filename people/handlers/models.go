package handlers

// Person .
type Person struct {
	ID        string   `json:"id"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Status    string   `json:"status"`
	Address   *Address `json:"address"`
}

// Address .
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}
