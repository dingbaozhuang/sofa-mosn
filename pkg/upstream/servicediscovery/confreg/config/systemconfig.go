package config

import (
    "github.com/magiconair/properties"
    "gitlab.alipay-inc.com/afe/mosn/pkg/log"
    "fmt"
    "strings"
)

type SystemConfig struct {
    AntShareCloud    bool
    InstanceId       string
    DataCenter       string
    AppName          string
    Zone             string
    RegistryEndpoint string
    AccessKey        string
    SecretKey        string
}

var SysConfig *SystemConfig

func InitSystemConfig(antShareCloud bool, dc string, appName string, zone string) *SystemConfig {
    if SysConfig != nil {
        return SysConfig
    }

    SysConfig = &SystemConfig{
        AntShareCloud: antShareCloud,
        DataCenter:    dc,
        AppName:       appName,
        Zone:          zone,
        InstanceId:    "DEFAULT_INSTANCE_ID",
    }

    confregUrl, z := readPropertyFromServerConfFile(antShareCloud)
    if SysConfig.Zone == "" {
        SysConfig.Zone = z
    }
    if !strings.HasPrefix(confregUrl, "http://") {
        confregUrl = "http://" + confregUrl
    }
    SysConfig.RegistryEndpoint = confregUrl
    return SysConfig
}

var serverConfFilePath = "/Users/lepdou/server.conf"

func readPropertyFromServerConfFile(antShareCloud bool) (confregUrl string, zone string) {
    if !antShareCloud {
        serverConf := properties.MustLoadFile(serverConfFilePath, properties.UTF8)
        cu, ok := serverConf.Get("confregurl")
        if !ok {
            errMsg := fmt.Sprintf("Load confregurl from %s failed.", serverConf)
            log.DefaultLogger.Errorf(errMsg)
            panic(errMsg)
        }
        z, _ := serverConf.Get("zone")
        return cu, z
    }
    return "", ""
}
