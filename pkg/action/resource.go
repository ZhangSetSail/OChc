package action

import (
	"fmt"
	"github.com/ZhangSetSail/OChc/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
)

type Resource struct {
	Namespace string
}

type ResourceOpt func(resource *Resource)

func NewResource() *Resource {
	p := &Resource{}
	return p
}

func (r *Resource) Run(args []string) {
	chaosPods := k8s.KCli.ListPods()
	chaosNodeIP := make(map[string]int)
	for _, p := range chaosPods {
		chaosNodeIP[p.Status.HostIP] = 1
	}
	nodes := k8s.KCli.ListNodes()
	for _, n := range nodes {
		address := n.Status.Addresses
		for _, addr := range address {
			if addr.Type == corev1.NodeInternalIP && chaosNodeIP[addr.Address] == 1 {
				fmt.Printf("addr %v ... %v\n", n.Status.NodeInfo.Architecture, addr.Address)
			}
		}
	}
}
