package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testConsulConfig = api.Config{
	Address: "http://192.168.116.91:8500",
}

func TestListServer(t *testing.T) {
	client, err := NewConsul(&testConsulConfig)
	assert.Equal(t, nil, err)

	services, err := client.ListService()
	assert.Equal(t, nil, err)

	fmt.Println(services)
}

func TestCheck(t *testing.T) {
	client, err := NewConsul(&testConsulConfig)
	assert.Equal(t, nil, err)

	exsit := client.Check()
	assert.Equal(t, true, exsit)
}

func TestListService(t *testing.T) {
	client, err := NewConsul(&testConsulConfig)
	assert.Equal(t, nil, err)

	services, err := client.ListService()
	assert.Equal(t, nil, err)

	for _, v := range services {
		_, err := client.GetServiceInfo(v.Name)
		assert.Equal(t, nil, err)

		//检查服务状态
		//_, err = client.CheckService(v.Name)
		//assert.NotEqual(t, nil, err)
	}
}

func TestKeySetAndSelect(t *testing.T) {
	client, err := NewConsul(&testConsulConfig)
	assert.Equal(t, nil, err)

	err = client.SetValue("test-key-1", []byte("test-value-1"))
	assert.NoError(t, err)

	value1, err := client.GetValue("test-key-1", nil)
	assert.NoError(t, err)
	assert.Equal(t, "test-value-1", string(value1))
}
