package helper

import (
	"fmt"

	"github.com/janritter/terrastate/backend/types"
)

func (helper *Helper) PrintStateFileAttributes(attributes []*types.StateFileAttribute) {
	fmt.Println("")
	fmt.Println("------- Using the following values -------")
	for _, attribute := range attributes {
		if attribute.Given {
			fmt.Printf("%s = %s \n", attribute.AttributeKey, attribute.Value)
		}
	}
	fmt.Println("------------------------------------------")
}
