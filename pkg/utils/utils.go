package utils

import (
	"errors"
	"github.com/Mangatsu/server/pkg/log"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

// ReadJSON returns the given JSON file as bytes.
func ReadJSON(jsonFile string) ([]byte, error) {
	jsonFileBytes, err := os.ReadFile(jsonFile)

	if err != nil {
		log.Z.Debug("failed to read JSON file", zap.String("err", err.Error()))
		return nil, err
	}

	return jsonFileBytes, nil
}

// Clamp clamps the given value to the given range.
func Clamp(value, min, max int64) int64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// PeriodicTask loops the given function in separate thread between the given interval.
func PeriodicTask(d time.Duration, f func()) {
	go func() {
		for {
			f()
			time.Sleep(d)
		}
	}()
}

// PathExists checks if the given path exists.
func PathExists(pathTo string) bool {
	_, err := os.OpenFile(pathTo, os.O_RDONLY, 0444)
	return !errors.Is(err, os.ErrNotExist)
}

func FileSize(filePath string) (int64, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		log.Z.Error("could not get file size", zap.String("path", filePath), zap.String("err", err.Error()))
		return 0, err
	}

	return stat.Size(), nil
}

func DirSize(dirPath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// Similarity calculates the similarity between two strings.
func Similarity(a string, b string) float64 {
	sd := metrics.NewSorensenDice()
	sd.CaseSensitive = false
	sd.NgramSize = 4
	similarity := strutil.Similarity(a, b, sd)

	return similarity
}
