package config

type Config struct {
	Profiles []Profile `yaml:"profiles"`
}

type Profile struct {
	Name             string `yaml:"name"`
	ConnectionString string `yaml:"connectionString"`
}

type CommandLineArgs struct {
	ConfigPath       string `short:"c" long:"config" description:"A path for configuration"`
	ConnectionString string `short:"s" long:"connection-string" description:"Connection string to use to connect to database"`
	Profile          string `short:"p" long:"profile" description:"A profile to use to connect to target database"`
	Query            string `short:"q" long:"query" description:"A query to execute"`
}
