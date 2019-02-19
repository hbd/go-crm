package people

import "context"

// DB .
type DB interface {
	PingDB() error

	GetPerson(context.Context, string) (Person, error)
	DeletePerson(context.Context, string) error
	UpsertPerson(context.Context, Person) error

	GetPeople(context.Context) ([]Person, error)
}
