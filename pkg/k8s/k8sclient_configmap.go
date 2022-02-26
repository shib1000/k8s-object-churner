package k8s

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k8sclient *K8sClient) GetConfigMaps(ns string, labelName string, labelValue string) (*v1.ConfigMapList, error) {
	configMaps, err := k8sclient.clientset.CoreV1().ConfigMaps(ns).
		List(context.TODO(),
			metav1.ListOptions{LabelSelector: fmt.Sprintf("%s=%s", labelName, labelValue)})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(configMaps.Items) == 0 {
		return nil, nil
	} else {
		return configMaps, nil
	}
}

func (k8sclient *K8sClient) CreateConfigMap(ns string,
	configMapName string,
	values map[string]string,
	labels map[string]string) (*v1.ConfigMap, error) {
	if labels == nil {
		labels = make(map[string]string)
	}
	
	labels["koc-created"] = "true"
	configMap := &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        configMapName,
			Namespace:   ns,
			Labels:      labels,
			Annotations: map[string]string{"koc-sample-object": "true"},
		},
		Data: values,
	}
	configMap, err := k8sclient.clientset.CoreV1().ConfigMaps(ns).Create(context.TODO(), configMap, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return configMap, nil
}

func (k8sclient *K8sClient) DeleteConfigMap(ns string, configMapName string) error {
	err := k8sclient.clientset.CoreV1().ConfigMaps(ns).Delete(context.TODO(), configMapName, metav1.DeleteOptions{})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
