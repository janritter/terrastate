package iface

type BackendAPI interface {
	GenerateStatefileForBackend(in interface{}) error
}
