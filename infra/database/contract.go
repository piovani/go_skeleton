package database

type DatabaseContract interface {
	Ping() error
}
