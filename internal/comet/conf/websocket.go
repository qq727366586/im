package conf

type Websocket struct {
	Bind        []string `yaml:"bind"`
	TLSOpen     bool     `yaml:"TLSOpen"`
	TLSBind     []string `yaml:"TLSBind"`
	CertFile    string   `yaml:"CertFile "`
	PrivateFile string   `yaml:"PrivateFile"`
}
