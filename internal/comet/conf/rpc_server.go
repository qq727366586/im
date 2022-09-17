package conf

type RPCServer struct {
	Network           string `yaml:"network"`
	Addr              string `yaml:"addr"`
	Timeout           int    `yaml:"timeout"`
	IdleTimeout       int    `yaml:"idleTimeout"`
	MaxLifeTime       int    `yaml:"maxLifeTime"`
	ForceCloseWait    int    `yaml:"forceCloseWait"`
	KeepAliveInterval int    `yaml:"keepAliveInterval"`
	KeepAliveTimeout  int    `yaml:"keepAliveTimeout"`
}
