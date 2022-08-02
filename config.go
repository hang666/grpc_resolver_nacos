package grpc_resolver_nacos

import (
	"net/url"

	"github.com/pkg/errors"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
)

type NacosClient = naming_client.INamingClient

type NacosInstance = model.Instance
type NacosService = model.Service

type NacosConfig struct {
	Label              string //用做寻找配置的标签
	ServiceName        string //标记服务名称
	Clusters           []string
	GroupName          string
	NacosServerConfigs []constant.ServerConfig
	NacosClientConfig  constant.ClientConfig
	NacosClient        NacosClient
}

var NacosConfigs map[string]*NacosConfig = make(map[string]*NacosConfig)

func AddNacosConfig(config NacosConfig) {
	//NacosConfigs.Store(config.Label, &config)
	NacosConfigs[config.Label] = &config
}

func DelNacosConfig(label string) {
	//NacosConfigs.Delete(label)
	delete(NacosConfigs, label)
}

func parseURL(u string) (*NacosConfig, error) {
	rawURL, err := url.Parse(u)
	if err != nil {
		return &NacosConfig{}, errors.Wrap(err, "Wrong nacos URL")
	}
	if rawURL.Scheme != schemeName || len(rawURL.Host) == 0 {
		return &NacosConfig{}, errors.Wrap(err, "Wrong nacos URL")
	}
	config, ok := NacosConfigs[rawURL.Host]
	if !ok {
		return &NacosConfig{}, errors.Wrap(err, "The nacos config is not exist")
	}
	return config, nil
}
