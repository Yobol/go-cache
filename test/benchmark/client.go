package main

type Cmd struct {
	Name  string
	Key   string
	Value string
	Error error
}

type ClientType string

const (
	ClientTypeHTTP = "http"
)

type Client interface {
	Run(*Cmd)
	PipelinedRun([]*Cmd)
}

func New(clientType ClientType) Client {
	switch clientType {
	case ClientTypeHTTP:
		return newHTTPClient()
	}
	panic("unknown server type ")
}
