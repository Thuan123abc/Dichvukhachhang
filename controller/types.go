package controller

type Phongban struct {
	ID   uint64
	Name string
}

type User struct {
	ID      uint64
	Phone   string
	Comment string
}

type Ticket struct {
	ID             uint64
	IDReception    uint64
	Phone          string
	Comment        string
	Content        string
	Status         bool
	TimeCompletion uint64
}
