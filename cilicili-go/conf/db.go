package conf

type Database struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}
