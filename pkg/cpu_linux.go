// +build !darwin

package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func Testing() {
	t := time.Now()
	ctimes, _ := cpu.Times(false)
	mycpu, _ := cpu.Info()
	fmt.Println(mycpu)
	fmt.Println(ctimes)
	fmt.Println(time.Since(t))

}
