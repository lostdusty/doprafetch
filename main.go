package main

import (
"fmt"
"log"
"os"
"runtime"
"os/exec"
"math"
"github.com/cloudfoundry/gosigar"
. "github.com/klauspost/cpuid"
"github.com/fatih/color"
//"github.com/printfcoder/goutils/intutils"
)

func main() {
red := color.New(color.FgRed).PrintfFunc() //print strings in red
underline := color.New(color.Bold, color.FgBlue).PrintfFunc()
info := color.New(color.Bold, color.FgRed).PrintfFunc()
errortxt := "DropaFetch failed to run correctly. Aborting.\nDetails: command failed with %s\n"
linuxonly := "Error: DropaLinux is meant to run only on Linux."


mem := sigar.Mem{}
mem.Get() //get memory
//ARMDetect() //get cpu


red("DropaFetch v1\n")
underline("- - - - - - -\n\n")
if runtime.GOOS != "linux" { color.Set(color.BgRed); log.Fatalf(linuxonly); color.Unset() }

info("OS: ")
fmt.Print("Dropa Linux ") //so far, this seems to be the name of the OS
cmd := exec.Command("uname", "-m")
cmd.Stdout = os.Stdout
err := cmd.Run()
if err != nil { log.Fatalf(errortxt, err) }


info("Host: ")
cmd = exec.Command("hostname")
cmd.Stdout = os.Stdout
err = cmd.Run()
if err != nil { log.Fatalf(errortxt, err) }

info("Kernel: ")
cmd = exec.Command("uname", "-r")
cmd.Stdout = os.Stdout
err = cmd.Run()
if err != nil { log.Fatalf(errortxt, err) }

info("Uptime: ") //TODO: Change this on future
cmd = exec.Command("uptime")
cmd.Stdout = os.Stdout
err = cmd.Run()
if err != nil { log.Fatalf(errortxt, err) }

info("Shell: ") //TODO: Fix this. I didn't found a consistent way to show busybox version, running the command above usually gives status code 127 or 2
/*( = exec.Command("ls", "--help", "2>&1", "&", "head", "-1", `|`, `cut`, `-f1`, "-d'('")
cmd.Stdout = os.Stdout
err = cmd.Run()
if err != nil { log.Fatalf(errortxt, err) }*/
fmt.Print("BusyBox\n")

info("CPU: ") 
fmt.Print(CPU.BrandName, "(", CPU.LogicalCores, ") @", CPU.Hz, "Hz\n") //Until now, only arm cpu flags is supported and theres nothing I can do about it. You still can use: cat /proc/cpuinfo

info("Memory: ")
mtotal := ToFloat64(mem.Total)
mused := ToFloat64(mem.Used)

fmt.Print(math.Trunc((mused / 2048) / 1000), " / ", math.Trunc((mtotal / 2048) / 1000), " MB\n")
}
//More information to come. Stay tuned!
func ToFloat64(in uint64) float64 {
	return float64(in)
}