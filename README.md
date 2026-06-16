# 🚀 Waesta Concurrent Downloader

> **Author:** waesta.js | **Platform:** Go

## 🇬🇧 English

### Overview
Multi-threaded HTTP Range downloader that splits files into concurrent chunks for maximum speed.

### ✨ Key Features
- **Modular Go Package:** `waesta.go` + `downloader.go`
- **WaestaLogger:** Banner and timestamped log functions
- **CLI Flags:** `-url`, `-chunks` for flexible configuration
- **Download Stats:** Tracks completed/failed chunks

### 🔒 Anti-Tamper Security
Integrity check at startup via `AntiTamperCheck()`.

### 📁 Project Structure
```
09_Waesta_Concurrent_Downloader/
├── main.go
├── waesta.go
└── downloader.go
```

### 🛠️ Usage
```bash
go run . -url https://example.com/file.zip
go run . -url https://example.com/file.zip -chunks 16
```

---

## 🇹🇷 Türkçe

### Genel Bakış
HTTP Range desteği ile dosyayı eşzamanlı parçalara bölerek indiren çok iş parçacıklı Go indirici.

### 🛠️ Kullanım
```bash
go run . -url https://example.com/dosya.zip -chunks 8
```
