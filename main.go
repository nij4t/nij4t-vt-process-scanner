package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

func getRawProcessList() ([]byte, error) {
	cmd := exec.Command("tasklist.exe", "/fo", "csv", "/nh")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return out, nil
}

func getProcessList() ([]string, error) {
	rawProc, err := getRawProcessList()

	proc := strings.Split(string(rawProc), "\n")

	procNames := []string{}
	for _, v := range proc {
		procNames = append(procNames, strings.Split(v, ",")[0])
	}

	if err != nil {
		return nil, err
	}

	return procNames, nil
}

func getProcessPath(str string) (string, error) {
	cmd := exec.Command(
		// "wmic",
		// "process",
		// "where",
		// "'name="+str+"'",
		// "get",
		// "ExecutablePath")
		"wmic", "process", "get", "ExecutablePath")
	// "process where 'name=\"chrome.exe\"' get ExecutablePath")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func filter(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		entry = strings.TrimSpace(entry)
		if _, v := keys[entry]; !v && strings.HasSuffix(entry, ".exe") {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func getRawProcessPaths() ([]byte, error) {
	cmd := exec.Command("wmic", "process", "get", "ExecutablePath")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return out, nil
}

func getProcessPaths() ([]string, error) {
	rawProc, err := getRawProcessPaths()

	proc := filter(strings.Split(string(rawProc), "\n"))

	// procNames := []string{}
	// for _, v := range proc {
	// 	procNames = append(procNames, strings.Split(v, "\n")[0])
	// }

	if err != nil {
		return nil, err
	}

	return proc, nil
}

func main() {
	// taskList, err := getProcessList()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// proc, err := getProcessPath(taskList[0])
	proc, err := getProcessPaths()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(strings.Join(taskList, "\n"))
	fmt.Println(strings.Join(proc, "\n"))
}
