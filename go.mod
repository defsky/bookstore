module github.com/defsky/bookstore

go 1.13

replace github.com/defsky/bookstore => ./

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.7.0
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.14
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
