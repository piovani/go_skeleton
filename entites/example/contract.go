package example

type ExampleRepository interface {
	Find(ID int) (*Example, error)
}
