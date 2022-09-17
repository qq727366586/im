package conf

type RPCClient struct {
	Dial    int `yaml:"dial"`
	Timeout int `yaml:"timeout"`
}
