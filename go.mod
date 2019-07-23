module github.com/inwinstack/blended

go 1.12

require (
	github.com/brancz/gojsontoyaml v0.0.0-20190425155809-e8bd32d46b3d // indirect
	github.com/emicklei/go-restful v2.9.6+incompatible // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/jsonnet-bundler/jsonnet-bundler v0.1.0 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20181027024800-9fcf73cc980b
	k8s.io/apimachinery v0.0.0-20180621070125-103fd098999d
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/code-generator v0.0.0-20190717022600-77f3a1fe56bb
)

replace (
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20181025213731-e84da0312774
	golang.org/x/net => golang.org/x/net v0.0.0-20190206173232-65e2d4e15006
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/text => golang.org/x/text v0.3.1-0.20181227161524-e6919f6577db
	golang.org/x/time => golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
	k8s.io/api => k8s.io/api v0.0.0-20181027024800-9fcf73cc980b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20180621070125-103fd098999d
	k8s.io/client-go => k8s.io/client-go v8.0.0+incompatible
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190717022600-77f3a1fe56bb
)
