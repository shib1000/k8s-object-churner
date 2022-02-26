package k8s

import (
	"context"
	"fmt"
	"github.com/shib1000/k8s-object-churner/pkg/azure"
	"github.com/shib1000/k8s-object-churner/pkg/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	clientset *kubernetes.Clientset
}

func NewK8sClient(appCfg *config.ConfigMgr) (k8sclient *K8sClient, err error) {
	azure.Init()
	var k8scfg *rest.Config
	kubeclientPath := appCfg.GetConfigString("KUBECONFIG")
	if kubeclientPath == "" {
		k8scfg, err = rest.InClusterConfig()
	} else {
		k8scfg, err = clientcmd.BuildConfigFromFlags("", kubeclientPath)
	}
	if err != nil {
		return nil, err
	}

	c, err := kubernetes.NewForConfig(k8scfg)
	if err != nil {
		return nil, err
	}
	return &K8sClient{clientset: c}, err
}

func (k8sclient *K8sClient) GetPodDetails(ns string, podname string) (*v1.Pod, error) {
	pods, err := k8sclient.clientset.CoreV1().Pods(ns).
		List(context.TODO(),
			metav1.ListOptions{FieldSelector: fmt.Sprintf("metadata.name=%s", podname)})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(pods.Items) == 0 {
		return nil, nil
	} else {
		return &pods.Items[0], nil
	}
}
