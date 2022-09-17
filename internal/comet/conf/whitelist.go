package conf

type Whitelist struct {
	List []int64 `yaml:"list"`
	Log  string  `yaml:"log"`
}
