package configo

import "testing"

func TestNewConfig(t *testing.T) {
	t.Log(NewConfig("/", TYPE_DEFAULT) != nil)
	t.Log(NewConfig("/") != nil)
}

func TestNewDefaultConfig(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Log("load error", err)
	}

}

func TestConfig_Get(t *testing.T) {
	t.Log(Get("some"))
}
