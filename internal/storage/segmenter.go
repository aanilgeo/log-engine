package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type LogStore struct {
	mu           sync.Mutex
	activeFile   *os.File
	currentSize  int64
	maxSize      int64 // e.g., 10 * 1024 * 1024 for 10MB
	dataDir      string
}

func NewLogStore(dir string, maxSize int64) (*LogStore, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	
	ls := &LogStore{
		dataDir: dir,
		maxSize: maxSize,
	}
	
	if err := ls.rotate(); err != nil {
		return nil, err
	}
	return ls, nil
}

func (ls *LogStore) Write(data []byte) error {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	// Check if we need to rotate before writing
	if ls.currentSize+int64(len(data)) > ls.maxSize {
		if err := ls.rotate(); err != nil {
			return err
		}
	}

	n, err := ls.activeFile.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	ls.currentSize += int64(n)
	return nil
}

func (ls *LogStore) rotate() error {
	if ls.activeFile != nil {
		ls.activeFile.Close()
	}

	fileName := fmt.Sprintf("%s/log_%d.seg", ls.dataDir, time.Now().UnixNano())
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	ls.activeFile = file
	ls.currentSize = 0
	return nil
}

func (ls *LogStore) Search(appID string, level string) ([]map[string]interface{}, error) {
	files, err := os.ReadDir(ls.dataDir)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for _, file := range files {
		if file.IsDir() { continue }
		
		f, err := os.Open(ls.dataDir + "/" + file.Name())
		if err != nil { continue }
		
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var entry map[string]interface{}
			if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
				continue
			}

			matchesApp := (appID == "" || entry["app_id"] == appID)
			matchesLevel := (level == "" || entry["level"] == level)

			if matchesApp && matchesLevel {
				results = append(results, entry)
			}
		}
		f.Close()
	}
	return results, nil
}