package database

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Ping() error {
	return nil
}
