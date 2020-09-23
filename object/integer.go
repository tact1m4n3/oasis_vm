package object

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) String() string {
	return fmt.Sprint(i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER
}
