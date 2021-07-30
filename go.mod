module github/fengjunhua/devops-backend-gin

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.2.0
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/prometheus/client_golang v1.10.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/inf.v0 v0.9.1 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
	k8s.io/api v0.0.0-20190111032252-67edc246be36
	k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93
	k8s.io/client-go v0.0.0-20190111032708-6bf63545bd02
	k8s.io/klog v1.0.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0
)

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
