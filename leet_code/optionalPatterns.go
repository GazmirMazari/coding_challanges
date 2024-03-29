package leet_code

type Server struct {
	Port int
	Host string
}

// OptionFunc is a type for functions that configure the ConfigurableService.
type OptionFunc func(*Server)

func withPort(port int) OptionFunc {
	return func(s *Server) {
		s.Port = port
	}
}

func withHost(host string) OptionFunc {
	return func(s *Server) {
		s.Host = host
	}
}

func NewConfigurableService(options ...OptionFunc) *Server {
	service := &Server{
		Host: "localhost",
		Port: 8080,
	}

	for _, option := range options {
		option(service)
	}

	return service
}

func main() {
	svr := NewConfigurableService(withPort(8081), withHost(""))
}
