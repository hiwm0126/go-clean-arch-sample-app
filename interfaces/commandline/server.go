package commandline

type Server interface {
	Run(args [][]string) error
}

type server struct {
	router Router
}

func NewServer() Server {
	return &server{
		router: NewRouter(),
	}
}

func (s *server) Run(args [][]string) error {
	err := s.router.Routing(args)
	if err != nil {
		return err
	}
	return nil
}
