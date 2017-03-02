package configo

import "testing"

func TestNewConfig(t *testing.T) {
	t.Log("ok")
}

func TestNewDefaultConfig(t *testing.T) {
	if err := config.Load(); err != nil {
		t.Log("load error", err)
	}

}
