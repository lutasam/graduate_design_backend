package utils

import (
	"errors"
	"path"
	"runtime"
	"sync"

	"github.com/spf13/viper"
)

type ConfigResolve struct {
	Viper *viper.Viper
}

var (
	Resolve     *ConfigResolve
	ResolveOnce sync.Once
)

func getConfigResolve() *ConfigResolve {
	ResolveOnce.Do(func() {
		Resolve = &ConfigResolve{
			Viper: viper.New(),
		}
		// achieve current filepath
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic(errors.New("config path cannot init"))
		}
		Resolve.Viper.SetConfigFile(path.Dir(filename) + "/../../conf/config.yml")
		Resolve.Viper.SetConfigType("yml")
		err := Resolve.Viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	})
	return Resolve
}

func GetConfigString(target string) string {
	return getConfigResolve().Viper.GetString(target)
}

func GetConfigIntSlice(target string) []int {
	return getConfigResolve().Viper.GetIntSlice(target)
}

func GetConfigStringSlice(target string) []string {
	return getConfigResolve().Viper.GetStringSlice(target)
}
