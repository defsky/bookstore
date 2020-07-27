package micro

// Option ...
type Option interface {
	String() string
	Value() interface{}
}

type nameOpt struct {
	v string
}

func (o nameOpt) String() string {
	return o.v
}
func (o nameOpt) Value() interface{} {
	return o.v
}

type versionOpt struct {
	v string
}

func (o versionOpt) String() string {
	return o.v
}
func (o versionOpt) Value() interface{} {
	return o.v
}

type registryOpt struct {
	v Registry
}

func (o registryOpt) String() string {
	return o.v.String()
}
func (o registryOpt) Value() interface{} {
	return o.v
}

// Name return name option for service
func Name(s string) Option {
	return nameOpt{
		v: s,
	}
}

// Version return version option for service
func Version(s string) Option {
	return versionOpt{
		v: s,
	}
}

// RegistryCenter return registry option for service
func RegistryCenter(r Registry) Option {
	return &registryOpt{v: r}
}
