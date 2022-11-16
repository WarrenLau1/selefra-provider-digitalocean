package digitalocean_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	Token             string   `yaml:"token"  mapstructure:"token"`
	SpacesRegions     []string `yaml:"spaces_regions"  mapstructure:"spaces_regions"`
	SpacesAccessKey   string   `yaml:"spaces_access_key"  mapstructure:"spaces_access_key"`
	SpacesAccessKeyId string   `yaml:"spaces_access_key_id"  mapstructure:"spaces_access_key_id"`
	AccountName       string   `yaml:"account_name,omitempty" mapstructure:"account_name"`
}
