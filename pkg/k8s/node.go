package k8s

import (
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *CliK8S) ListNodes() []corev1.Node {
	nodes, err := k.clientSet.CoreV1().Nodes().List(k.ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Errorf("list nodes failure: %v", err)
	}
	return nodes.Items
}
