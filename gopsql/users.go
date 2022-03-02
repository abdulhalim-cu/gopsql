package gopsql

type User struct {
	ID       int
	Username string
}

type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}
