package endpoints

import "example-service/internal/usecases"

type Endpoints struct {
	Usecases *usecases.Usecases
}

func NewEndpoints(usecases *usecases.Usecases) *Endpoints {
	return &Endpoints{Usecases: usecases}
}

func (e *Endpoints) Foo() {

}
