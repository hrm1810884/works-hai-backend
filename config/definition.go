package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//REF: https://qiita.com/takehanKosuke/items/1b17ade882b50cf2d737

// マッピング用の構造体
type Config struct {
	Server   Server   `yaml:"server"`
	Firebase Firebase `yaml:"firbase"`
	Python   Python   `yaml:"python"`
}

type Server struct {
	Dev string `yaml:"dev"`
	Api string `yaml:"api"`
}

type Firebase struct {
	Bucket string `yaml:"bucket"`
}

type Python struct {
	Host     string `yaml:"host"`
	Endpoint string `yaml:"endpoint"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")          // 設定ファイル名を指定
	viper.SetConfigType("yaml")            // 設定ファイルの形式を指定
	viper.AddConfigPath("config/private/") // ファイルのpathを指定

	err := viper.ReadInConfig() // 設定ファイルを探索して読み取る
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %w", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &cfg, nil
}
