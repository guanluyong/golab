package gpa

import (
	"fmt"
)

type IGpaRegister interface {
	RegisterEntity(entity interface{});
}

type IGpaFactory interface {
	Save(entity interface{});
}