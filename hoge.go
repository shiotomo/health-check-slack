package main


import (
	"fmt"
	"github.com/shiotomo/health-check-slack/pkg/utils"
)

func main() {
	fmt.Println(utils.GetHostIpAddr())
}