package database

type BodyPart struct {
	Id   int
	Name string
}

func GetAllBodyPart() (b []*BodyPart, err error) {
	err = db.Select(b, "SELECT * FROM body_part")
	return b, err
}

func GetByIdBodyPart(id int) (b *BodyPart, err error) {
	err = db.Get(b, "SELECT * FROM body_part WHERE id = $1", id)
	return b, err
}
