package creator

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/janritter/terrastate/backend/types"
)

type CreatorAPI interface {
	Create([]*types.StateFileAttribute)
}

type Creator struct {
}

func NewCreator() *Creator {
	return &Creator{}
}

func (creator *Creator) Create(stateFileAttributes []*types.StateFileAttribute, backend string) {
	f, err := os.Create("terrastate.tf")
	defer f.Close()

	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	_, err = f.WriteString("terraform {\n  backend \"" + backend + "\" {\n")
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	for _, attribute := range stateFileAttributes {
		if attribute.Given {
			value := fmt.Sprintf("%v", attribute.Value)
			_, err = f.WriteString("		" + attribute.AttributeKey + " = \"" + value + "\"\n")
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
		}
	}
	_, err = f.WriteString("  }\n}")
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	color.Green("Successfully created terrastate.tf")

}
