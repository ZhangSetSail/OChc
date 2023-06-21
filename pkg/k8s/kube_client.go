package k8s

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
	"time"
)

type CliK8S struct {
	clientSet *kubernetes.Clientset
	Namespace string
	ctx       context.Context
	cancel    context.CancelFunc
}

func InitK8SCli() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Fatal(err)
	}
	// 创建 client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx = ctx
	KCli = &CliK8S{
		clientSet: clientSet,
		ctx:       ctx,
		cancel:    cancel,
	}
}

var KCli *CliK8S
