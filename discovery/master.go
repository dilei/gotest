package discovery

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

type Master struct {
	members map[string]*Member
	KeysAPI *clientv3.Client
}

// Member is a clientv3 machine
type Member struct {
	InGroup bool
	IP      string
	Name    string
	CPU     int
}

func NewMaster(endpoints []string) *Master {

	// 为了保证 HTTPS 链接可信，需要预先加载目标证书签发机构的 CA 根证书
	etcdCA, err := ioutil.ReadFile("D:/goproj/src/gotest/discovery/etcd-root-ca.pem")
	if err != nil {
		log.Fatal(err)
	}

	// etcd 启用了双向 TLS 认证，所以客户端证书同样需要加载
	etcdClientCert, err := tls.LoadX509KeyPair("D:/goproj/src/gotest/discovery/s1.pem", "D:/goproj/src/gotest/discovery/s1-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个空的 CA Pool
	// 因为后续只会链接 Etcd 的 api 端点，所以此处选择使用空的 CA Pool，然后只加入 Etcd CA 既可
	// 如果期望链接其他 TLS 端点，那么最好使用 x509.SystemCertPool() 方法先 copy 一份系统根 CA
	// 然后再向这个 Pool 中添加自定义 CA
	rootCertPool := x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM(etcdCA)

	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
		// 自定义 CA 及 Client Cert 配置
		TLS: &tls.Config{
			RootCAs:      rootCertPool,
			Certificates: []tls.Certificate{etcdClientCert},
		},
	}

	etcdclientv3, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	master := &Master{
		members: make(map[string]*Member),
		KeysAPI: etcdclientv3,
	}
	go master.WatchWorkers()
	return master
}

func (m *Master) AddWorker(info *WorkerInfo) {
	member := &Member{
		InGroup: true,
		IP:      info.IP,
		Name:    info.Name,
		CPU:     info.CPU,
	}
	m.members[member.Name] = member
}

func (m *Master) UpdateWorker(info *WorkerInfo) {
	member := m.members[info.Name]
	member.InGroup = true
}

func (m *Master) WatchWorkers() {
	api := m.KeysAPI
	rch := api.Watch(context.TODO(), "workers/")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
