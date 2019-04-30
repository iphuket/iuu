package cpu

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/shirou/gopsutil/cpu"
)

// Info ...cpu info json
func Info(c *gin.Context) {

	ci, err := cpu.Info()
	if err != nil {
		c.JSON(404, gin.H{"errInfo": fmt.Sprint(err)})
		c.Abort()
	}

	cu, err := cpu.ProcInfo()
	if err != nil {
		c.JSON(404, gin.H{"errInfo": fmt.Sprint(err)})
		c.Abort()
	}

	cpc, err := cpu.Percent(time.Second, true)
	if err != nil {
		c.JSON(404, gin.H{"errInfo": fmt.Sprint(err)})
		c.Abort()
	}
	fmt.Println("cu", len(cu))

	c.JSON(200, gin.H{"Cores": ci[0].Cores, "Ghz": ci[0].Mhz / 1000, "ModeName": ci[0].ModelName, "Processes": cu[0].Processes, "Percent": cpc})

}

// Usage for cpu messages
func Usage(c *gin.Context) {
	cu, err := cpu.ProcInfo()
	if err != nil {
		fmt.Println("cu", err)
	}
	fmt.Println(cu[0].Processes)            // cpu 进程数
	fmt.Println(cu[0].ProcessorQueueLength) // CPU 列队长度
	fmt.Println(cpu.Percent(time.Second, false))
}
