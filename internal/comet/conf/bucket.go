package conf

type Bucket struct {
	Size          int    `yaml:"size"`
	Channel       int    `yaml:"channel"`
	Room          int    `yaml:"room"`
	RoutineAmount uint64 `yaml:"routineAmount"`
	RoutineSize   int    `yaml:"routineSize"`
}
