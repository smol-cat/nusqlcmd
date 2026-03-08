package config

type Config struct {
	Profiles []Profile
}

type Profile struct {
	Name             string `yaml:"name"`
	ConnectionString string `yaml:"connectionString"`
}
