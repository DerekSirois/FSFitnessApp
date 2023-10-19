package database

type Muscle struct {
	Id   int
	Name string
}

func GetAllMuscle() (b []*Muscle, err error) {
	err = db.Select(&b, "SELECT * FROM muscle")
	return b, err
}

func GetByIdMuscle(id int) (*Muscle, error) {
	b := &Muscle{}
	err := db.Get(b, "SELECT * FROM muscle WHERE id = $1", id)
	return b, err
}
