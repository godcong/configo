package configo

import (
	"os"
	"runtime"
	"strings"
)

type CONFIG_TYPE int

const (
	TYPE_DEFAULT CONFIG_TYPE = iota
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
	Type      CONFIG_TYPE
	Path      string
	Configure Common
}

var config *Config

const DEFAULT_CONFIG_FILE = "config.env"

func init() {
	config = NewDefaultConfig()
	if e := config.Load(); e != nil {
		config.Configure = make(Default)
	}
}

//GetSystemSeparator System Separator
func GetSystemSeparator() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

//NewDefaultConfig default config
func NewDefaultConfig() *Config {
	wd, err := os.Getwd()
	fp := ""
	if err == nil {
		fp = strings.Join([]string{wd, DEFAULT_CONFIG_FILE}, GetSystemSeparator())
	}

	return NewConfig(fp)
}

//NewConfig new config
func NewConfig(path string, args ...CONFIG_TYPE) *Config {
	defaultConfig := make(Default)

	configType := TYPE_DEFAULT
	if args != nil {
		configType = args[0]
	}

	conf := &Config{
		Type:      configType,
		Path:      path,
		Configure: (Common)(defaultConfig),
	}
	return conf
}

//Load default load
func Load() error {
	return config.Load()
}

//Load load
func (c *Config) Load() error {
	file, openErr := os.Open(c.Path)
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
	if c.Type == TYPE_DEFAULT {
		return envDefault(c, f)
	}

	return nil
}

//Get default get
func Get(s string) (*Property, error) {
	return config.Get(s)
}

//Get config get
func (c *Config) Get(s string) (*Property, error) {
	if config.Type == TYPE_DEFAULT {
		p := envDefaultGet(s)
		if p != nil {
			return p, nil
		}
		return nil, ERROR_CONFIG_GET_PROPERTY
	}
	return nil, ERROR_CONFIG_GET_PROPERTY_TYPE
}

//Get property get
func (p *Property) Get(s string) (string, error) {

	if v, ok := (*p)[s]; ok {
		return v, nil
	}

	return "", ERROR_CONFIG_GET_PROPERTY_VALUE

}

//MustGet property get with default value
func (p *Property) MustGet(s, d string) string {
	if v, ok := (*p)[s]; ok {
		return v
	}
	return d
}
