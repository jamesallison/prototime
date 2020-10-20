package main

import (
	"fmt"
	"time"

	"github.com/monzo/wearedev/libraries/util"
)

func main() {
	fmt.Print(util.TimeToProto(time.Now().UTC()))
}
