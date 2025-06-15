package main
import (
		"fmt"

    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/mem"

)
type SysInfo struct {
    Hostname string `bson:hostname`
    Platform string `bson:platform`
    CPU      string `bson:cpu`
    RAM      uint64 `bson:ram`
    Disk     uint64 `bson:disk`
}
func main() {
    hostStat, _ := host.Info()
    cpuStat, _ := cpu.Info()
    vmStat, _ := mem.VirtualMemory()
    diskStat, _ := disk.Usage("/home/khat/Documents/")

    info := new(SysInfo)

    info.Hostname = hostStat.Hostname
    info.Platform = hostStat.Platform
    info.CPU = cpuStat[0].ModelName
    info.RAM = vmStat.Total / 1024 / 1024
    info.Disk = diskStat.Total / 1024 / 1024

		fmt.Println("Hostname: ",info.Hostname)
		fmt.Println("Platform: ",hostStat.Platform, hostStat.PlatformVersion)
		fmt.Println("CPU: ", cpuStat[0].ModelName)
		fmt.Println("RAM: ", info.RAM, " MB")
		fmt.Println("Disk: ", info.Disk," MB")
}


