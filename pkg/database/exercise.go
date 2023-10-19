package database

type Exercise struct {
	Id          int
	Name        string
	Description string
	MuscleId    int `db:"muscle_id"`
}

func GetAllExercises() (e []*Exercise, err error) {
	err = db.Select(&e, "SELECT * FROM exercise")
	return e, err
}

func GetByIdExercise(id int) (e *Exercise, err error) {
	err = db.Get(e, "SELECT * FROM exercise WHERE id = $1", id)
	return e, err
}

func CreateExercise(e *Exercise) error {
	_, err := db.Exec("INSERT INTO exercise(name, description, muscle_id) VALUES ($1, $2, $3)", e.Name, e.Description, e.MuscleId)
	return err
}

func UpdateExercise(e *Exercise) error {
	_, err := db.Exec("UPDATE exercise SET name = $1, description = $2, muscle_id = $3 WHERE id = $4", e.Name, e.Description, e.MuscleId, e.Id)
	return err
}

func DeleteExercise(id int) error {
	_, err := db.Exec("DELETE FROM exercise WHERE id = $1", id)
	return err
}
