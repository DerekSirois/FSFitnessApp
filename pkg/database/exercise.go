package database

type Exercise struct {
	Id          int
	Name        string
	Description string
	BodyPartId  int
}

func GetAllExercises() (e []*Exercise, err error) {
	err = db.Select(&e, "SELECT * FROM exercise")
	return e, err
}

func GetByIdExercise(id int) (e *Exercise, err error) {
	err = db.Get(e, "SELECT * FROM exercise WHERE id = $1", id)
	return e, err
}

func Create(e *Exercise) error {
	_, err := db.Exec("INSERT INTO exercise(name, description, body_part_id) VALUES ($1, $2, $3)", e.Name, e.Description, e.BodyPartId)
	return err
}

func Update(e *Exercise) error {
	_, err := db.Exec("UPDATE exercise SET name = $1, description = $2, body_part_id = $3 WHERE id = $4", e.Name, e.Description, e.BodyPartId, e.Id)
	return err
}

func Delete(id int) error {
	_, err := db.Exec("DELETE FROM exercise WHERE id = $1", id)
	return err
}
