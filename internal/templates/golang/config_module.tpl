package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"os"
)

var Module = fx.Provide(
	RegisterAppConfigs,
)

func RegisterAppConfigs() *viper.Viper {
	configsFiles := []string{
		"config/config.yaml",
		"/etc/secrets/config.yaml",
	}
	var filename string

	for _, file := range configsFiles {
		_, err := os.Stat(file)
		if errors.Is(err, os.ErrNotExist) {
            continue
        }
		filename = file
	}
	if filename == "" {
		log.Fatalf("config files not found")
		os.Exit(1)
	}

	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error read config: %v", err)
		return nil
	}
	return viper.GetViper()
}
