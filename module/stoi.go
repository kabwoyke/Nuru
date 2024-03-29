package module

import (
	"strconv"

	"github.com/AvicennaJr/Nuru/object"
)

var StoiFunction = map[string]object.ModuleFunction{}

func init() {
	StoiFunction["tungo_to_numbari"] = tungo_to_numbari
}

func tungo_to_numbari(args []object.Object, defs map[string]object.Object) object.Object {

	if len(defs) != 0 {
		return &object.Error{Message: "Hoja hii hairuhusiwi"}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Tunahitaji hoja moja tu"}
	}
	if args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Hoja lazima iwe neno"}
	}
	input := args[0].(*object.String).Value
	intValue, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return &object.Error{Message: "Oparesheni ya kubadilisha aina string to number imefeli"}
	}
	return &object.Integer{Value: intValue}
}
