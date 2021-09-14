package process

type User struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Email     string   `json:"email,omitempty"`
	Status    string   `json:"status,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	State string `json:"state,omitempty"`
	City  string `json:"city,omitempty"`
}
