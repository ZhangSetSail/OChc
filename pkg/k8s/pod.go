package k8s

import (
	"fmt"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *CliK8S) ListPods() []corev1.Pod {
	fmt.Println(k.Namespace)
	pods, err := k.clientSet.CoreV1().Pods(k.Namespace).List(k.ctx, metav1.ListOptions{
		LabelSelector: "name=rbd-chaos",
	})
	if err != nil {
		logrus.Errorf("list nodes failure: %v", err)
	}
	return pods.Items
}
