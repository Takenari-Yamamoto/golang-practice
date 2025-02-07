package main

import "fmt"

// Configはオプションフィールドを持つ構造体
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

// ConfigOptionはオプションフィールドを設定する関数型
type ConfigOption func(*Config)

// NewConfigは必須のHostとPortを受け取り、オプションフィールドを適用するコンストラクタ
func NewConfig(host string, port int, opts ...ConfigOption) *Config {
	config := &Config{
		Host: host,
		Port: port,
	}

	for _, opt := range opts {
		opt(config)
	}

	return config
}

// WithUsernameはオプションフィールドを設定するためのヘルパー関数
func WithUsername(username string) ConfigOption {
	return func(c *Config) {
		c.Username = username
	}
}

// WithPasswordはオプションフィールドを設定するためのヘルパー関数
func WithPassword(password string) ConfigOption {
	return func(c *Config) {
		c.Password = password
	}
}

func main() {
	config := NewConfig("localhost", 8080, WithUsername("admin"), WithPassword("secret"))

	fmt.Println(config)
}
