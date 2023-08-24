package config

type Mongo struct {
	Uri string
	Db  string
}

type Config struct {
	Mongo Mongo `yaml:"Mongo"`
}
