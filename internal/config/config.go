package config

type LogConfig struct {
	Level string `yaml:"level" json:"level"`
}

type StorageConfig struct {
	DigestAlgorithm string `yaml:"digest-algorithm" json:"digest-algorithm"`
}

type AuthConfig struct {
	Password string `yaml:"password" json:"password"`
}

type TLSConfig struct {
	Enabled *bool `yaml:"enabled" json:"enabled"`
}

type NetworkConfig struct {
	Address string `yaml:"host" json:"host"`
}

type Config struct {
	Log     LogConfig     `yaml:"log" json:"log"`
	Storage StorageConfig `yaml:"storage" json:"storage"`
	Auth    AuthConfig    `yaml:"auth" json:"auth"`
	TLS     TLSConfig     `yaml:"tls" json:"tls"`
	Network NetworkConfig `yaml:"network" json:"network"`
}
