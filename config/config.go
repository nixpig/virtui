package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	DatabaseKey = "database"
	LogfileKey  = "log"

	filename = "virtui.toml"
	database = "virtui.db"
	logfile  = "virtui.log"
)

func Initialise(path string, v *viper.Viper) error {
	if path != "" {
		_, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("check provided config path exists: %w", err)
		}
	} else {
		userConfigDir, err := os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("get user config dir: %w", err)
		}

		path = filepath.Join(userConfigDir, filename)
	}

	configFile, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_CREATE,
		0666,
	)
	if err != nil {
		return fmt.Errorf("open config file (%s): %w", path, err)
	}
	defer configFile.Close()

	v.SetConfigFile(configFile.Name())
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read in config: %w", err)
	}

	if v.GetString(DatabaseKey) == "" {
		v.Set(DatabaseKey, database)
	}

	if v.GetString(LogfileKey) == "" {
		v.Set(LogfileKey, logfile)
	}

	if err := v.WriteConfig(); err != nil {
		return fmt.Errorf("write default config: %w", err)
	}

	return nil
}
