package conf

import (
	"fmt"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	err := LoadConfigFromEnv()

	if err != nil {
		t.Fatal(err, "env parse error")
	}

	//
	//t.Log(C())
	fmt.Println(C(), "config")
}
