package kvs

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

/*type UserClean func(use User)

func (u *User) toDomain() *process.User {

	user := &process.User{
		ID: u.ID, FirstName: u.FirstName, LastName: u.LastName, Status: u.Status}

	return user
}
func userFromDomain(userDom process.User) *User {
	u := User{}
	u.ID = userDom.ID
	u.FirstName = userDom.FirstName
	u.LastName = userDom.LastName
	u.Status $userDom.Status
	return &u
}*/
