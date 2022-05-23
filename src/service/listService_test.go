package service

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestListService(t *testing.T) {
	stdout, _ := exec.Command("go", "env", "GOMOD").Output()
	path := string(bytes.TrimSpace(stdout))
	if path == "" {
		os.Exit(1)
	}
	ss := strings.Split(path, "\\")
	ss = ss[:len(ss)-1]
	path = strings.Join(ss, "\\") + "\\"
	log.Println(path)
}
