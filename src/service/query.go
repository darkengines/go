package service

type Metadata struct {
	ServiceName string
}

type Query struct {
	Metadata Metadata
	Parameter interface{}
}