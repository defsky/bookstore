package micro

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

// Client ...
type Client struct {
	Name, Version string

	r Registry

	c *grpc.ClientConn
}

// NewClient ...
func NewClient(opts ...Option) *Client {
	c := &Client{
		Version: "latest",
		r:       defaultRegistry,
	}
	if drv := os.Getenv("MICRO_REGISTRY"); len(drv) > 0 {
		log.Println("registry driver: etcd")
		c.r = NewRegistry(drv)
	}
	for _, o := range opts {
		switch o.(type) {
		case nameOpt:
			c.Name = o.String()
		case versionOpt:
			c.Version = o.String()
		case registryOpt:
			c.r = o.Value().(Registry)
		}
	}

	return c
}

// Conn is *grpc.ClientConn
func (c *Client) Conn() *grpc.ClientConn {
	if c.c != nil {
		return c.c
	}
	if len(c.Name) <= 0 {
		log.Fatalln("must specify service name")
	}
	serverAddr, err := c.r.Query(c)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	c.c = conn

	return c.c
}
