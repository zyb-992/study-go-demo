package config

type Etcd struct {
	EndPoint string `json:"end_point"`
}

func (c *Config) EtcdConf() {

}
