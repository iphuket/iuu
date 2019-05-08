package cpu

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"

	"github.com/shirou/gopsutil/cpu"
)

// Info ...cpu info json
func Info(c *gin.Context) {
	cp, err :=cpu.Percent(time.Duration(1/1000)*time.Second,true)
	if err !=nil {
		c.JSON(c.Writer.Status(), gin.H{"err": fmt.Sprint(err)})
	}
	c.JSON(c.Writer.Status(), cp)
}

// Usage for cpu messages
func Usage(c *gin.Context) {
	/*
		cu, err := cpu.ProcInfo()
		if err != nil {
			fmt.Println("cu", err)
		}
		fmt.Println(cu[0].Processes)            // cpu 进程数
		fmt.Println(cu[0].ProcessorQueueLength) // CPU 列队长度
		fmt.Println(cpu.Percent(time.Second, false))
	*/
}
