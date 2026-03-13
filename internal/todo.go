package internal

type ToDo interface {
	Foo() (err error)
}

type ToDoImpl struct{}

func (t *ToDoImpl) Foo() error {
	return nil
}
