package object

import "fmt"

type ObjectType string

const (
	INTEGER ObjectType = "integer"
)

type Object interface {
	fmt.Stringer
	Type() ObjectType
}
