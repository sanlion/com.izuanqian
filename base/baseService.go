package base

import (
	"fmt"
	"github.com/sluu99/uuid"
	"strings"
)

func SayHello(name string) {

	fmt.Println(name + " come in.[" + strings.ToUpper(strings.Replace(uuid.Rand().Hex(), "-", "", -1)) + "]")

}
