// +build darwin

package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func Testing() error {
	t := time.Now()
	mycpu, err := cpu.Info()
	fmt.Println(mycpu)
	fmt.Println("THis is the darwin build")
	fmt.Println(time.Since(t))
	return err

}
