package client

type Cmd struct {
	Name  string
	Key   string
	Value string
	Error error
}

type ClientType string

const (
	ClientTypeHTTP = "http"
	ClientTypeTCP  = "tcp"
)

type Client interface {
	Run(*Cmd)
	PipelinedRun([]*Cmd)
}

func New(clientType ClientType, remoteAddr string) Client {
	switch clientType {
	case ClientTypeHTTP:
		return newHTTPClient()
	case ClientTypeTCP:
		return newTCPClient(remoteAddr)
	}
	panic("unknown server type " + clientType)
}
