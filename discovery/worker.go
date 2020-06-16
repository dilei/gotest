package discovery

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

type Worker struct {
	Name    string
	IP      string
	KeysAPI *clientv3.Client
}

// workerInfo is the service register information to etcd
type WorkerInfo struct {
	Name string
	IP   string
	CPU  int
}

func NewWorker(name, IP string, endpoints []string) *Worker {
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

	w := &Worker{
		Name:    name,
		IP:      IP,
		KeysAPI: etcdclientv3,
	}
	return w
}

func (w *Worker) HeartBeat() {
	api := w.KeysAPI
	for {
		info := &WorkerInfo{
			Name: w.Name,
			IP:   w.IP,
			CPU:  runtime.NumCPU(),
		}

		key := "workers/" + w.Name
		value, _ := json.Marshal(info)

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		_, err := api.Put(timeoutCtx, key, string(value))
		if err != nil {
			log.Println("Error update workerInfo:", err)
		}

		timeoutCtx2, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		// op example
		// clientv3.WithPrefix()
		res, err := api.Get(timeoutCtx2, key, clientv3.WithPrefix())
		if err != nil {
			log.Println("Error:", err)
		}
		fmt.Println(string(res.Kvs[0].Value))
		time.Sleep(time.Second * 2)
	}
}
