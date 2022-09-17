package conf

type Protocol struct {
	Timer            int `yaml:"timer"`
	TimerSize        int `yaml:"timerSize"`
	SvrProto         int `yaml:"svrProto"`
	CliProto         int `yaml:"cliProto"`
	HandshakeTimeout int `yaml:"handshakeTimeout"` // 秒
}
