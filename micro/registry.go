package micro

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	etcd "github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var defaultRegistry = NewRegistry("mdns")

// Registry is registry center interface for service
type Registry interface {
	Register(*Service)
	UnRegister(*Service)
	Query(*Client) (string, error)
	String() string
}

// NewRegistry return a registry center object
func NewRegistry(driver string) Registry {
	switch driver {
	case "etcd":
		if a := os.Getenv("MICRO_REGISTRY_ADDRESS"); len(a) > 0 {
			log.Printf("registry address: %s\n", a)
			cfg := etcd.Config{
				Endpoints:   strings.Split(a, ","),
				DialTimeout: 5 * time.Second,
			}
			return &etcdRegistry{
				address: a,
				config:  cfg,
				close:   make(chan struct{}),
			}
		}
		log.Fatalln("please specify registry address with MICRO_REGISTRY_ADDRESS")
		return nil

	case "mdns":
		return &mdnsRegistry{}
	}

	log.Fatalln(fmt.Sprintf("undefined registry driver: %s", driver))
	return nil
}

type etcdRegistry struct {
	config  etcd.Config
	address string
	close   chan struct{}
}

type serviceInfo struct {
	// IPs ip addresses
	UUID string   `json:"UUID"`
	IPs  []string `json:"ips"`
	Port int      `json:"port"`
}

func (r *etcdRegistry) String() string {
	return fmt.Sprintf("etcd %s", r.address)
}
func (r *etcdRegistry) Register(s *Service) {
	log.Printf("register service [etcd] %s:%s:%s\n", s.Name, s.Version, s.UUID)
	ip := make([]string, 0)
	port := 0
	if a, ok := s.l.Addr().(*net.TCPAddr); ok && !a.IP.IsUnspecified() && a.IP.To4() != nil {
		port = a.Port
		ip = append(ip, a.IP.String())
	} else {
		port = a.Port
		ips, err := getHostIPs()
		if err == nil {
			ip = append(ip, ips...)
		}
	}
	if len(ip) > 0 {

	} else {
		log.Fatalln("no valid ip address")
	}
	if port == 0 {
		log.Fatalln("invalid port number")
	}

	etcdClient, err := etcd.New(r.config)
	if err != nil {
		log.Fatalln(err)
	}
	lease := etcd.NewLease(etcdClient)
	leaseResp, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatalln(err)
	}
	leaseID := leaseResp.ID

	ctx, cancelFunc := context.WithCancel(context.TODO())
	lkar, err := lease.KeepAlive(ctx, leaseID)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
	DONE:
		for {
			select {
			case resp := <-lkar:
				if resp == nil {
					cancelFunc()
					goto DONE
				}
			case <-r.close:
				cancelFunc()
				goto DONE
			}
		}
	}()

	kv := etcd.NewKV(etcdClient)

	info := &serviceInfo{
		IPs:  ip,
		Port: port,
	}
	infoJSON, err := json.Marshal(info)
	if err != nil {
		log.Println(err)
	}
	_, err = kv.Put(context.TODO(),
		fmt.Sprintf("/micro/%s/%s/%s", s.Name, s.Version, s.UUID),
		string(infoJSON),
		etcd.WithLease(leaseID))
	if err != nil {
		log.Println(err)
	}
}

func (r *etcdRegistry) UnRegister(s *Service) {
	log.Printf("unregister service [etcd] %s:%s:%s\n", s.Name, s.Version, s.UUID)
	r.close <- struct{}{}
}
func (r *etcdRegistry) Query(c *Client) (string, error) {
	log.Printf("query service [etcd] %s:%s\n", c.Name, c.Version)

	etcdClient, err := etcd.New(r.config)
	if err != nil {
		log.Fatalln(err)
	}
	kv := etcd.NewKV(etcdClient)
	resp, err := kv.Get(context.TODO(), fmt.Sprintf("/micro/%s/%s", c.Name, c.Version), etcd.WithPrefix())
	if err != nil {
		log.Fatalln(err)
	}
	var info *serviceInfo

	// query instances for service at specify version
DONE:
	for _, v := range resp.Kvs {

		svcInfo := &serviceInfo{}
		if err := json.Unmarshal(v.Value, svcInfo); err == nil {
			k := strings.Split(string(v.Key), "/")
			sid := k[len(k)-1]
			svcInfo.UUID = sid
			info = svcInfo
			break DONE
		} else {
			log.Println(err)
		}
	}

	// TODO: load balance between multiple instances

	log.Printf("query response: %v\n", info)
	if info != nil && len(info.IPs) > 0 {
		return fmt.Sprintf("%s:%d", info.IPs[0], info.Port), nil
	}
	return "", status.Errorf(codes.NotFound, fmt.Sprintf("service %s not found", c.Name))
}

type mdnsRegistry struct {
	domain string
}

func (r *mdnsRegistry) String() string {
	return fmt.Sprintf("mdns %s", r.domain)
}
func (r *mdnsRegistry) Register(s *Service) {
	log.Printf("register service [mdns] %s:%s, address %s\n", s.Name, s.Version, s.l.Addr())
}
func (r *mdnsRegistry) UnRegister(s *Service) {
	log.Printf("unregister service [mdns] %s:%s, address %s\n", s.Name, s.Version, s.l.Addr())
}
func (r *mdnsRegistry) Query(c *Client) (string, error) {
	log.Printf("query service [mdns] %s:%s\n", c.Name, c.Version)

	switch c.Name {
	case "com.afkplayer.service.user:latest":
		return "user:50051", nil
	}

	return "", status.Errorf(codes.NotFound, fmt.Sprintf("service %s not found", c.Name))
}

func getHostIPs() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	ips := make([]string, 0)
	for _, a := range addrs {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
			ips = append(ips, ip.IP.String())
		}
	}
	if len(ips) > 0 {
		return ips, nil
	}

	return nil, errors.New("no invalid ip address")
}
