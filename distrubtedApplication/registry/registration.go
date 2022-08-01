package registry

type Resgistration struct {
	ServiceName ServiceName
	ServiceURL  string
}

type ServiceName string

const (
	LogService = ServiceName("LogService")
)
