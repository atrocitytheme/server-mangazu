package library

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Mangatsu/server/internal/config"
	"github.com/Mangatsu/server/pkg/db"
	"github.com/Mangatsu/server/pkg/log"
	"go.uber.org/zap"
)

func walk(libraryPath string, libraryID int32, libraryLayout config.Layout) fs.WalkDirFunc {
	return func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		isImage := ImageExtensions.MatchString(d.Name())
		isArchive := ArchiveExtensions.MatchString(d.Name())
		if !isArchive && !isImage {
			return nil
		}

		s = filepath.ToSlash(s)
		relativePath := config.RelativePath(libraryPath, s)

		// If an image is found, the parent dir will be considered as a gallery.
		if isImage {
			relativePath = path.Dir(relativePath)
		}

		// Skip if already in database
		if db.ArchivePathFound(relativePath) {
			log.Z.Debug("skipping archive already in db", zap.String("name", d.Name()))
			return nil
		}

		// Series name from the dir name if Structured layout
		series := ""
		if libraryLayout == config.Structured {
			dirs := strings.SplitN(relativePath, "/", 2)
			series = dirs[0]
		}

		var title string
		if isImage {
			title = path.Base(relativePath)
		} else {
			n := strings.LastIndex(d.Name(), path.Ext(d.Name()))
			title = d.Name()[:n]
		}

		uuid, err := db.NewGallery(relativePath, libraryID, title, series)
		if err != nil {
			log.Z.Error("failed to add gallery to db",
				zap.String("path", relativePath),
				zap.String("err", err.Error()))
		} else {
			// Generates cover thumbnail
			go ReadArchiveImages(config.BuildLibraryPath(libraryPath, relativePath), uuid, true)
			log.Z.Debug("added gallery", zap.String("path", relativePath))
		}

		if isImage {
			return fs.SkipDir
		}

		return err
	}
}

func ScanArchives() {
	// TODO: Quick scan by only checking directories that have been modified.
	// Not too important as current implementation is pretty fast already.
	libraries, err := db.GetOnlyLibraries()
	if err != nil {
		log.Z.Error("failed to find libraries to scan", zap.String("err", err.Error()))
		return
	}

	for _, library := range libraries {
		err := filepath.WalkDir(library.Path, walk(library.Path, library.ID, config.Layout(library.Layout)))
		if err != nil {
			log.Z.Error("skipping library as an error occurred during scanning",
				zap.String("path", library.Path),
				zap.String("err", err.Error()))
			continue
		}
	}
}

func PathExists(pathTo string) bool {
	_, err := os.OpenFile(pathTo, os.O_RDONLY, 0444)
	return !errors.Is(err, os.ErrNotExist)
}
