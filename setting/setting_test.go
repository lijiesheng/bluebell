package setting

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init("../conf/config.yaml")
	fmt.Println(Conf)
}
