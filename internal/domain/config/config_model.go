package config

type Mongo struct {
	Uri string `yaml:"uri"`
	Db  string `yaml:"db"`
}

type Aws struct {
	AccessKey       string `yaml:"accessKey"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	Region          string `yaml:"region"`
}

type Config struct {
	Mongo Mongo `yaml:"Mongo"`
	Aws   Aws   `yaml:"Aws"`
}
