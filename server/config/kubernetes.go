package config

type Kubernetes struct {
	Kubeconfig string `mapstructure:"kubeconfig" json:"kubeconfig" yaml:"kubeconfig"` // path-to-kubeconfig
	InCluster  bool   `mapstructure:"in-cluster" json:"in-cluster" yaml:"in-cluster"` // 是否运行中pod中
}
