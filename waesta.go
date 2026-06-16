package main

import (
	"fmt"
	"os"
	"time"
)

const Author = "waesta.js"
const ExpectedHash = 927

func AntiTamperCheck() {
	sum := 0
	for _, char := range Author {
		sum += int(char)
	}
	if sum != ExpectedHash {
		fmt.Println("[FATAL] Waesta Integrity Check Failed.")
		os.Exit(1)
	}
}

func Banner(title, subtitle string) {
	fmt.Printf(`
========================================
 %s
 Developer: %s
 %s
========================================

`, title, Author, subtitle)
}

func Info(message string) {
	fmt.Printf("[INFO] %s %s\n", time.Now().Format("15:04:05"), message)
}

func Success(message string) {
	fmt.Printf("[ OK ] %s %s\n", time.Now().Format("15:04:05"), message)
}

func Warn(message string) {
	fmt.Printf("[WARN] %s %s\n", time.Now().Format("15:04:05"), message)
}

func Error(message string) {
	fmt.Printf("[ERR ] %s %s\n", time.Now().Format("15:04:05"), message)
}
