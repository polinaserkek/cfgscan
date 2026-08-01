package config

type LogConfig struct {
	Level string `yaml:"level" json:"level"`
}

type StorageConfig struct {
	DigestAlgorithm string `yaml:"digest-algorithm" json:"digest-algorithm"`
}

type Config struct {
	Log     LogConfig     `yaml:"log" json:"log"`
	Storage StorageConfig `yaml:"storage" json:"storage"`
}
