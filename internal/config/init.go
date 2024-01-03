package config

import (
	"github.com/spf13/viper"
	"go-backend/internal/util"
)

// MainConfig is the main config struct
type MainConfig struct {
	Database Database `json:"database"`
	Logging  Logging  `json:"logging"`
	Server   Server   `json:"server"`
	Features Features `json:"features"`
}

// Config MainConfig
var Config MainConfig

func InitConfig() {
	sugar := util.Logger.Sugar()
	// loading config
	sugar.Info("Loading config...")
	// Read in environment-specific config if it exists
	viper.AutomaticEnv()          // read in environment variables that match
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("json")   // config file type
	sugar.Info("ENV MODE:", viper.GetString("APP_ENV"))
	viper.AddConfigPath(
		viper.GetString("CONFIG_PATH"),
	) // path to config files
	// load config type
	env := viper.GetString("APP_ENV")
	if env == "production" {
		viper.SetConfigName("config.prod")
	} else {
		viper.SetConfigName("config.dev")
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		sugar.Info("Using config file:", viper.ConfigFileUsed())
		// unmarshal config
		errUnmarshal := viper.Unmarshal(&Config)
		if errUnmarshal != nil {
			sugar.Fatal("Error unmarshalling config file:", err)
		}
	} else {
		sugar.Fatal("Error loading config file:", err)
	}
}
