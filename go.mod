module gotest

go 1.12

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.44.3
	cloud.google.com/go/datastore => cloud.google.com/go/datastore v1.0.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190826190057-c7b8b68b1456
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200522201501-cb1345f3a375
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543

)

require (
	github.com/astaxie/session v0.0.0-20130408050157-95d7fe18579c
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/daizuozhuo/etcd-service-discovery v0.0.0-20170813133207-32e3e614e2d5
	github.com/go-redis/redis/v7 v7.0.0-beta.5
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/philchia/agollo v2.1.0+incompatible
	github.com/valyala/fasthttp v1.4.0
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	golang.org/x/tools v0.0.0-20200522201501-cb1345f3a375 // indirect
	google.golang.org/grpc v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.2.5
)
