package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
	"github.com/fatih/color"
	. "github.com/klauspost/cpuid"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	//Strings.
	red := color.New(color.FgRed).PrintfFunc() //Print strings in red
	underline := color.New(color.Bold, color.FgBlue).PrintfFunc()
	info := color.New(color.Bold, color.FgRed).PrintfFunc()
	errortxt := "DropaFetch failed to run correctly. Aborting.\nDetails: command failed with %s\n"
	linuxonly := "Error: DropaLinux is meant to run only on Linux."

	//Get system memory & uptime.
	mem := sigar.Mem{}
	mem.Get() //get memory
	concreteSigar := sigar.ConcreteSigar{}
	uptime := sigar.Uptime{}
	uptime.Get()
	avg, err := concreteSigar.GetLoadAverage()
	if err != nil {
		fmt.Printf("Failed to get load average")
		return
	}

	//Start the application
	red("DropaFetch v1\n")
	underline("- - - - - - -\n\n")
	if runtime.GOOS != "linux" { //If the application is running outside a Linux system it will abort.
		color.Set(color.BgRed)
		log.Fatalf(linuxonly)
		color.Unset()
	}

	info("OS: ")
	fmt.Print("Dropa Linux ") //So far, this seems to be the name of the OS and theres no way to get it from uname or other way.
	cmd := exec.Command("uname", "-m")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("Host: ")
	cmd = exec.Command("hostname")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("Kernel: ")
	cmd = exec.Command("uname", "-r")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("Uptime: ") //TODO: Show only the uptime.
	fmt.Print(os.Stdout, " %s up %s load average: %.2f, %.2f, %.2f\n",
		time.Now().Format("15:04:05"),
		uptime.Format(),
		avg.One, avg.Five, avg.Fifteen)

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

	fmt.Print(math.Trunc((mused/2048)/1000), " / ", math.Trunc((mtotal/2048)/1000), " MB\n")
}

func ToFloat64(in uint64) float64 {
	return float64(in)
}

//More information to come. Stay tuned!
