package backend

import (
	"errors"
	"log"

	"github.com/janritter/terrastate/backend/iface"
	"github.com/janritter/terrastate/backend/s3"
)

func GetBackendInterface(backendType string) (iface.BackendAPI, error) {
	switch backendType {
	case "s3":
		return s3.NewS3Backend(), nil
	}

	err := errors.New("Backend " + backendType + " is currently not supported")
	log.Println(err)
	return nil, err
}
