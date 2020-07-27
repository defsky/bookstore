module github.com/defsky/bookstore

go 1.13

replace github.com/defsky/bookstore => ./

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.14
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.26.0
)
