package conf

type TCP struct {
	Bind         []string `yaml:"bind"`
	Sndbuf       int      `yaml:"sndbuf"`
	Rcvbuf       int      `yaml:"rcvbuf"`
	KeepAlive    bool     `yaml:"keepAlive"`
	Reader       int      `yaml:"reader"`
	ReadBuf      int      `yaml:"readBuf"`
	ReadBufSize  int      `yaml:"readBufSize"`
	Writer       int      `yaml:"writer"`
	WriteBuf     int      `yaml:"writeBuf"`
	WriteBufSize int      `yaml:"writeBufSize"`
}
