package limberr

import (
	"errors"
	"testing"
)

func TestAddCode(t *testing.T) {
	// tests := []struct  {
	// 	name string
	// }{
	// 	{
	// 		name: "test name",
	// 	},
	// }

	// for _, test := range tests {
	// 	t.Run(test.name, func(t *testing.T) {

	// 	})
	// }

	err := errors.New("first error need to be with code")
	err2 := AddCode(err, "E983212")

	t.Log(err2)
}
