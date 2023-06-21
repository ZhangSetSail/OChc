package k8s

import (
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *CliK8S) ListNamespace() []corev1.Namespace {
	namespaces, err := k.clientSet.CoreV1().Namespaces().List(k.ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Errorf("list namespaces failure: %v", err)
	}
	return namespaces.Items
}
