package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Parent interface {
	GetName() string
}

func GetName[T Parent](parameter T) string {
	return parameter.GetName()
}

type Mother interface {
	GetName() string
	getMotherName() string
}

type MyMother struct {
	Name string
}

func (m MyMother) GetName() string {
	return m.Name
}

func (m MyMother) getMotherName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "yuni", GetName[Mother](&MyMother{Name: "yuni"}))
}
