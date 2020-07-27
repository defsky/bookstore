package micro

import (
	"log"
	"net"
	"os"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var serviceAddress = ":50051"

// Service ...
type Service struct {
	// S is grpc server
	S *grpc.Server

	l net.Listener
	r Registry

	Name, Version, UUID string
}

// NewService ...
func NewService(opts ...Option) *Service {
	sid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}
	s := &Service{
		Version: "latest",
		r:       defaultRegistry,
		UUID:    sid.String(),
	}
	if drv := os.Getenv("MICRO_REGISTRY"); len(drv) > 0 {
		s.r = NewRegistry(drv)
	}

	for _, o := range opts {
		switch o.(type) {
		case nameOpt:
			s.Name = o.String()
		case versionOpt:
			s.Version = o.String()
		case registryOpt:
			s.r = o.Value().(Registry)
		}
	}

	return s
}

// Init ...
func (s *Service) Init() *Service {
	if len(s.Name) <= 0 {
		log.Fatalln("must specify service name")
	}
	if addr := os.Getenv("MICRO_SERVER_ADDRESS"); len(addr) > 0 {
		serviceAddress = addr
	}
	lis, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatalln(err)
	}
	s.l = lis
	s.S = grpc.NewServer()

	return s
}

// Serve ...
func (s *Service) Serve() error {
	log.Printf("service listening on %s\n", s.l.Addr())

	s.r.Register(s)

	err := s.S.Serve(s.l)

	s.r.UnRegister(s)

	return err
}
