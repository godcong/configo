package configo

import (
	"os"
	"runtime"
	"strings"
)

const (
	TYPE_DEFAULT = iota
)

const (
	PROCESS_NONE = iota
	PROCESS_COMMON
	PROCESS_PROPERTY
)

type (
	Common   interface{}
	Property map[string]string
	Default  map[string]Property
)

type Config struct {
	configType int
	configPath string
	configure  Common
}

var config *Config

const CONFIG_FILE = "config.env"

func init() {
	config = NewDefaultConfig()
	if e := config.Load(); e != nil {
		panic(e)
	}
}

func GetSystemSeparator() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

func NewDefaultConfig() *Config {
	defaultConfig := make(Default)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fp := strings.Join([]string{wd, CONFIG_FILE}, GetSystemSeparator())
	conf := &Config{
		configType: TYPE_DEFAULT,
		configPath: fp,
		configure:  (Common)(defaultConfig),
	}
	return conf
}

func Load() error {
	return config.Load()
}

func (c *Config) Load() error {
	file, openErr := os.Open(c.configPath)
	if openErr != nil {
		return ERROR_CONFIG_CANNOT_OPEN
	}
	defer file.Close()

	if e := envLoad(c, file); e != nil {
		return e
	}
	return nil
}

func envLoad(c *Config, f *os.File) error {
	if c.configType == TYPE_DEFAULT {
		return envDefault(c, f)
	}

	return nil
}

func Get(s string) *Property {
	return config.Get(s)
}

func (c *Config) Get(s string) *Property {
	if config.configType == TYPE_DEFAULT {
		return envDefaultGet(s)
	}
	return nil
}

func (p *Property) Get(s ...string) string {
	if len(s) > 1 {
		return p.GetOneD(s[0], s[1])
	} else {
		return p.GetOne(s[0])
	}
}

func (p *Property) GetOne(s string) string {
	if v, ok := (*p)[s]; ok {
		return v
	}

	return ""
}

func (p *Property) GetOneD(s, d string) string {
	if v, ok := (*p)[s]; ok {
		return v
	}
	return d
}
