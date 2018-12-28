// +build !darwin

package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func Testing() error {
	t := time.Now()
	fmt.Println("This is a linux build")
	ctimes, _ := cpu.Times(false)
	mycpu, err := cpu.Info()
	fmt.Println(mycpu)
	fmt.Println(ctimes)
	fmt.Println(time.Since(t))

	return err
}
