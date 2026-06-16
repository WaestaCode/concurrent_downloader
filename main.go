package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	AntiTamperCheck()
	Banner("🚀 WAESTA CONCURRENT DOWNLOADER", "Multi-Thread HTTP Range Download Engine")

	url := flag.String("url", "", "İndirilecek dosya URL'si")
	chunks := flag.Int("chunks", 8, "Eşzamanlı parça sayısı")
	flag.Parse()

	if *url == "" && len(os.Args) > 1 && os.Args[1][0] != '-' {
		*url = os.Args[1]
	}

	if *url == "" {
		fmt.Println("Kullanım: go run . -url <URL>")
		fmt.Println("         go run . <URL>")
		os.Exit(1)
	}

	start := time.Now()
	if err := RunConcurrentDownload(*url, *chunks); err != nil {
		Error(err.Error())
		os.Exit(1)
	}

	Info(fmt.Sprintf("Toplam süre: %v | Secured by %s", time.Since(start).Round(time.Millisecond), Author))
}
