package store

// Interface ...
type Interface interface {
	Read(string) (interface{}, bool)
	Write(string, interface{})
}
