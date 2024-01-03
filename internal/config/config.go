package config

import "github.com/spf13/viper"

type Config struct {
	Server *ServerConfig `mapstructure:"server"`
	Image  *ImageConfig  `mapstructure:"image"`
}

type ServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type ImageConfig struct {
	DefaultFormat string       `mapstructure:"default_format"`
	Size          *SizeConfig  `mapstructure:"size"`
	Color         *ColorConfig `mapstructure:"color"`
}

type SizeConfig struct {
	MinSize   string `mapstructure:"min_size"`
	MaxSize   string `mapstructure:"max_size"`
	SizeClamp bool   `mapstructure:"size_clamp"`
}

type ColorConfig struct {
	DefaultBackgroundColor string `mapstructure:"default_background_color"`
	DefaultTextColor       string `mapstructure:"default_text_color"`
}

func LoadConfig(fileName string) (*viper.Viper, error) {
	cfgFile := viper.New()

	cfgFile.AddConfigPath("config")
	cfgFile.SetConfigName(fileName)
	cfgFile.AutomaticEnv()

	if err := cfgFile.ReadInConfig(); err != nil {
		return nil, err
	}

	return cfgFile, nil
}

func ParseConfig(cfgFile *viper.Viper) (*Config, error) {
	cfg := &Config{}

	if err := cfgFile.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
