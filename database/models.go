package database

type PhongbanModel struct {
	ID   uint64
	Name string
}

func (p *PhongbanModel) TableName() string {
	return "phongbans"
}

type UserModel struct {
	ID      uint64
	Phone   string
	Comment string
}

func (u *UserModel) TableName() string {
	return "users"
}

type TicketModel struct {
	ID             uint64
	IDReception    uint64 `gorm:"foreignKey:Phongban.ID"`
	Phone          string
	Comment        string
	Content        string
	Status         bool
	TimeCompletion uint64
}

func (t *TicketModel) TableName() string {
	return "tickets"
}
