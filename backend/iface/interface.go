package iface

type BackendAPI interface {
	GenerateConfigurationForBackend(in interface{}) error
}
