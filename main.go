package main

import (
	"main/picscan"
	"os"
)

func main() {
	curDir, _ := os.Getwd()
	picscan.ScanAndRenamePics(curDir)
}
