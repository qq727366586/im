package conf

import (
	"flag"
	"github.com/bilibili/discovery/naming"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
)

var (
	confPath  string
	region    string
	zone      string
	deployEnv string
	host      string
	addrs     string
	weight    int64
	offline   bool
	debug     bool

	// Conf config
	Conf *Config
)

type Env struct {
	Region    string
	Zone      string
	DeployEnv string
	Host      string
	Weight    int64
	Offline   bool
	Addrs     []string
}

type Config struct {
	Debug     bool           `yaml:"debug"`
	Env       *Env           `yaml:"env"`
	Discovery *naming.Config `yaml:"discovery"`
	TCP       *TCP           `yaml:"TCP"`
	Websocket *Websocket     `yaml:"websocket"`
	Protocol  *Protocol      `yaml:"protocol"`
	Bucket    *Bucket        `yaml:"bucket"`
	RPCClient *RPCClient     `yaml:"rpc_client"`
	RPCServer *RPCServer     `yaml:"rpc_server"`
	Whitelist *Whitelist     `yaml:"whitelist"`
}

func inits() {
	var (
		defHost, _    = os.Hostname()
		defAddrs      = os.Getenv("ADDRS")
		defWeight, _  = strconv.ParseInt(os.Getenv("WEIGHT"), 10, 32)
		defOffline, _ = strconv.ParseBool(os.Getenv("OFFLINE"))
		defDebug, _   = strconv.ParseBool(os.Getenv("DEBUG"))
	)
	flag.StringVar(&confPath, "conf", "comet-example.toml", "default config path.")
	flag.StringVar(&region, "region", os.Getenv("REGION"), "avaliable region. or use REGION env variable, value: sh etc.")
	flag.StringVar(&zone, "zone", os.Getenv("ZONE"), "avaliable zone. or use ZONE env variable, value: sh001/sh002 etc.")
	flag.StringVar(&deployEnv, "deploy.env", os.Getenv("DEPLOY_ENV"), "deploy env. or use DEPLOY_ENV env variable, value: dev/fat1/uat/pre/prod etc.")
	flag.StringVar(&host, "host", defHost, "machine hostname. or use default machine hostname.")
	flag.StringVar(&addrs, "addrs", defAddrs, "server public ip addrs. or use ADDRS env variable, value: 127.0.0.1 etc.")
	flag.Int64Var(&weight, "weight", defWeight, "load balancing weight, or use WEIGHT env variable, value: 10 etc.")
	flag.BoolVar(&offline, "offline", defOffline, "server offline. or use OFFLINE env variable, value: true/false etc.")
	flag.BoolVar(&debug, "debug", defDebug, "server debug. or use DEBUG env variable, value: true/false etc.")
}

func Default() *Config {
	return &Config{
		Debug:     debug,
		Env:       &Env{Region: region, Zone: zone, DeployEnv: deployEnv, Host: host, Weight: weight, Addrs: strings.Split(addrs, ","), Offline: offline},
		Discovery: &naming.Config{Region: region, Zone: zone, Env: deployEnv, Host: host},
		RPCClient: &RPCClient{
			Dial:    1,
			Timeout: 1,
		},
		RPCServer: &RPCServer{
			Network:           "tcp",
			Addr:              ":3109",
			Timeout:           1,
			IdleTimeout:       60,
			MaxLifeTime:       2 * 60 * 60,
			ForceCloseWait:    20,
			KeepAliveInterval: 60,
			KeepAliveTimeout:  20,
		},
		TCP: &TCP{
			Bind:         []string{":3101"},
			Sndbuf:       4096,
			Rcvbuf:       4096,
			KeepAlive:    false,
			Reader:       32,
			ReadBuf:      1024,
			ReadBufSize:  8192,
			Writer:       32,
			WriteBuf:     1024,
			WriteBufSize: 8192,
		},
		Websocket: &Websocket{
			Bind: []string{":3102"},
		},
		Protocol: &Protocol{
			Timer:            32,
			TimerSize:        2048,
			CliProto:         5,
			SvrProto:         10,
			HandshakeTimeout: 5,
		},
		Bucket: &Bucket{
			Size:          32,
			Channel:       1024,
			Room:          1024,
			RoutineAmount: 32,
			RoutineSize:   1024,
		},
	}
}

func Init() (err error) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path)   //设置读取的文件路径
	config.SetConfigName("base") //设置读取的文件名
	config.SetConfigType("yaml") //设置文件的类型
	if err = config.ReadInConfig(); err != nil {
		panic(err)
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	return
}
