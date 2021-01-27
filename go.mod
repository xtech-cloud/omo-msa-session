module omo.msa.session

go 1.15

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.26.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/micro/go-micro/v2 v2.3.0
	github.com/micro/go-plugins/config/source/consul/v2 v2.0.3
	github.com/micro/go-plugins/logger/logrus/v2 v2.3.0
	github.com/micro/go-plugins/registry/consul/v2 v2.0.3
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.3
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/xtech-cloud/omo-msp-session v1.1.0
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.4
)
