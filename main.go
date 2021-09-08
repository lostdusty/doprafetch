package main

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
	"github.com/fatih/color"
	"github.com/ricochet2200/go-disk-usage/du"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
)

var KB = uint64(1024)

func main() {
	//Strings.
	red := color.New(color.FgRed).PrintfFunc() //Print strings in red
	underline := color.New(color.Bold, color.FgBlue).PrintfFunc()
	info := color.New(color.Bold, color.FgRed).PrintfFunc()
	errortxt := "DropaFetch failed to run correctly. Aborting.\nDetails: command failed with %s\n"

	//Get system memory
	mem := sigar.Mem{}
	mem.Get()

	//Start the application
	if runtime.GOOS != "linux" { //Make sure it's running under a linux system.
		color.Set(color.BgRed)
		fmt.Println("Error: DropaFetch should run only on Linux.")
		color.Unset()
		os.Exit(2)
	}

	red("DopraFetch v2\n")

	underline("- - - - - - -\n\n")

	info("OS: ")
	//So far, this seems to be the name of the OS and theres no way to get it from uname or other way.
	cmd := exec.Command("uname", "-mrs")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("Firmware: ") //
	cmd = exec.Command(`sh`, `-c`, "cat /etc/DL-Release")
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

	info("Uptime: ") //Big thanks to Tonkku for helping me out!
	cmd = exec.Command(`sh`, `-c`, "uptime | awk '{print $1}'")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("Shell: ")
	cmd = exec.Command("sh", "-c", "busybox | head -1 | cut -f1 -d'('")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("CPU:")
	cmd = exec.Command("sh", "-c", "cat /proc/cpuinfo | grep 'model name' | cut -f2 -d':'")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("CPU Count: ")
	cmd = exec.Command("sh", "-c", "cat /proc/cpuinfo | grep processor | wc -l")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalf(errortxt, err)
	}

	info("Disk Usage: ") //I won't be using gosigar here since it didn't detected any disks.
	usage := du.NewDiskUsage("/")
	usagecd := du.NewDiskUsage(".")
	fmt.Print(usage.Used()/(KB*KB), " / ", usage.Size()/(KB*KB), " MB (/) \n")
	fmt.Print(usagecd.Used()/(KB*KB), " / ", usagecd.Size()/(KB*KB), " MB (Current Dir) \n")

	info("Memory: ")
	mtotal := ToFloat64(mem.Total)
	mused := ToFloat64(mem.Used)

	fmt.Print(math.Trunc((mused/2048)/1000), " / ", math.Trunc((mtotal/2048)/1000), " MB\n")
}

func ToFloat64(in uint64) float64 {
	return float64(in)
}
