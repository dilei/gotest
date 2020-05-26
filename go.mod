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
	github.com/go-redis/redis/v7 v7.0.0-beta.5
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/mock v1.4.3
	github.com/philchia/agollo v2.1.0+incompatible
	github.com/valyala/fasthttp v1.4.0
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
	golang.org/x/tools v0.0.0-20200522201501-cb1345f3a375 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	gopkg.in/yaml.v2 v2.2.5
)
