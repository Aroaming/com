package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net/http"
)

type Consul struct {
	ConsulClient *api.Client
	ConsulConfig *api.Config
	//	Lc           logger.LoggingClient
	//	TokenProvide interfaces.SecretProviderExt
}

// new consul
func NewConsul(config *api.Config) (*Consul, error) {
	var err error
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Consul{
		ConsulClient: client,
		ConsulConfig: config,
	}, nil
}

func (c *Consul) RefreshToken(token string) error {
	var err error
	c.ConsulConfig.Token = token
	//重新生成客户端
	c.ConsulClient, err = api.NewClient(c.ConsulConfig)
	if err != nil {
		return err
	}
	return nil
}

// GetValueByKey 获取值通过key
func (c *Consul) GetValue(key string, query *api.QueryOptions) ([]byte, error) {
	// 创建键值对查询选项
	//options := &api.QueryOptions{}

	// 获取键值对
	kvPair, _, err := c.ConsulClient.KV().Get(key, query)
	//retry, err := c.reloadAccessTokenOnAuthError(err)
	//if retry {
	//	// Try again with new Access Token
	//	kvPair, _, err = c.ConsulClient.KV().Get(keyInfo, options)
	//}
	if err != nil {
		log.Printf("获取consul value失败,key:%s,error:%v", key, err.Error())
		return nil, err
	}
	if kvPair == nil {
		log.Printf("key：%s不存在", key)
		return nil, fmt.Errorf("没有该键的配置%s", key)
	}
	// 检查键值对是否存在
	return kvPair.Value, nil
}

// SetValueByKey 设置值通过key
func (c *Consul) SetValue(key string, b []byte) error {
	// 创建键值对
	kv := &api.KVPair{
		Key:   key,
		Value: b,
	}

	// 设置键值对
	_, err := c.ConsulClient.KV().Put(kv, nil)
	if err != nil {
		log.Printf("设置consul key:%s,失败", key)
		return err
	}
	return nil
}

// GetKeysByType 查询所有consul中指定目录的key
func (c *Consul) GetKeys(dir string) ([]string, error) {
	pairs, _, err := c.ConsulClient.KV().List(dir, nil)
	if err != nil {
		log.Printf("查询consul键列表失败：%s", dir)
		return nil, err
	}
	var keys []string
	for _, v := range pairs {
		keys = append(keys, string(v.Value))
	}
	return keys, nil
}

// GetService 获取单个服务信息
func (c *Consul) GetServiceInfo(name string) (Service, error) {
	serviceInfo, _, err := c.ConsulClient.Agent().Service(name, nil)
	if err != nil {
		log.Printf("获取服务:%s信息失败：%s", name, err.Error())
		return Service{}, err
	}

	return Service{
		Address: serviceInfo.Address,
		Port:    serviceInfo.Port,
	}, nil
}

// ListServer 获取已注册的服务
func (c *Consul) ListService() ([]Service, error) {
	services, err := c.ConsulClient.Agent().Services()
	if err != nil {
		log.Printf("获取服务列表失败：%s", err.Error())
		return nil, err
	}

	var moduleList []Service

	for _, service := range services {
		moduleList = append(moduleList, Service{
			Name:    service.Service,
			Address: service.Address,
			Port:    service.Port,
		})
	}

	return moduleList, nil
}

// IsAlive 检测consul是否运行
func (c *Consul) Check() bool {
	url := fmt.Sprintf("http://%s/%s", c.ConsulConfig.Address, consulStatusPath)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("无法连接到Consul API:%s", err.Error())
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}

// CheckServer 检测服务
func (c *Consul) CheckService(serviceName string) (bool, error) {
	status, _, err := c.ConsulClient.Agent().AgentHealthServiceByID(serviceName)
	if err != nil {
		return false, err
	}
	if status == serviceStatusPass {
		return true, nil
	}
	return false, fmt.Errorf("server status %s", status)
}
