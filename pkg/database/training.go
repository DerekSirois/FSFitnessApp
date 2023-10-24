package database

type Training struct {
	Id      int
	Name    string
	Weekday string
	UserId  int `db:"user_id"`
}

func GetAllTraining(userId int) (t []*Training, err error) {
	err = db.Select(&t, "SELECT id, name, weekday FROM training WHERE user_id = $1", userId)
	return t, err
}

func GetByIdTraining(id int) (*Training, error) {
	t := &Training{}
	err := db.Get(t, "SELECT id, name, weekday FROM training WHERE id = $1", id)
	return t, err
}

func CreateTraining(t *Training) error {
	_, err := db.Exec("INSERT INTO training (name, weekday, user_id) VALUES ($1, $2, $3)", t.Name, t.Weekday, t.UserId)
	return err
}

func UpdateTraining(t *Training) error {
	_, err := db.Exec("UPDATE training SET name = $1, weekday = $2 WHERE id = $3", t.Name, t.Weekday, t.Id)
	return err
}

func DeleteTraining(id int) error {
	_, err := db.Exec("DELETE FROM training WHERE id = $1", id)
	return err
}
