package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type DownloadStats struct {
	Completed int
	Failed    int
}

func downloadChunk(url string, start, end, chunkID int, wg *sync.WaitGroup, stats *DownloadStats) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		stats.Failed++
		Error(fmt.Sprintf("Chunk %d request error", chunkID))
		return
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	req.Header.Set("User-Agent", fmt.Sprintf("WaestaDownloader/1.0 (%s)", Author))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		stats.Failed++
		Error(fmt.Sprintf("Chunk %d failed", chunkID))
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(fmt.Sprintf("waesta_chunk_%d.tmp", chunkID))
	if err != nil {
		stats.Failed++
		Error(fmt.Sprintf("Chunk %d file error", chunkID))
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		stats.Failed++
		Error(fmt.Sprintf("Chunk %d write error", chunkID))
		return
	}

	stats.Completed++
	Success(fmt.Sprintf("Chunk %d indirildi", chunkID))
}

func RunConcurrentDownload(url string, chunks int) error {
	Info(fmt.Sprintf("Metadata alınıyor → %s", url))

	resp, err := http.Head(url)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("dosya boyutu alınamadı veya sunucu Range desteklemiyor")
	}
	defer resp.Body.Close()

	size := int(resp.ContentLength)
	if size <= 0 {
		return fmt.Errorf("geçersiz dosya boyutu")
	}

	Info(fmt.Sprintf("Dosya boyutu: %d bytes | Parça: %d", size, chunks))

	chunkSize := size / chunks
	var wg sync.WaitGroup
	stats := &DownloadStats{}

	for i := 0; i < chunks; i++ {
		start := i * chunkSize
		end := start + chunkSize - 1
		if i == chunks-1 {
			end = size - 1
		}
		wg.Add(1)
		go downloadChunk(url, start, end, i, &wg, stats)
	}

	wg.Wait()

	if stats.Failed > 0 {
		Warn(fmt.Sprintf("%d parça başarısız, %d parça tamamlandı", stats.Failed, stats.Completed))
	} else {
		Success(fmt.Sprintf("Tüm parçalar indirildi → waesta_chunk_*.tmp"))
	}

	return nil
}
