package rpc

type Registry interface {
	Registry() error
	Deregistry() error
}
