package sql

import (
	"fmt"
	"github.com/pipe-cd/pipecd/pkg/plugin/sdk"
	"os"
	"path/filepath"
	"regexp"
)

// ExtractArtifactVersions extracts version information from SQL migration scripts.
func ExtractArtifactVersions(scriptPath string) ([]*sdk.ArtifactVersion, error) {
	if scriptPath == "" {
		return nil, fmt.Errorf("scriptPath must not be empty")
	}
	// Assume scriptPath is a single file or directory of migration scripts.
	info, err := os.Stat(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat script path %s: %v", scriptPath, err)
	}
	var files []string
	if info.IsDir() {
		entries, err := os.ReadDir(scriptPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read directory %s: %v", scriptPath, err)
		}
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".sql" {
				files = append(files, filepath.Join(scriptPath, entry.Name()))
			}
		}
	} else {
		files = []string{scriptPath}
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no SQL files found at %s", scriptPath)
	}
	// Regular expression to match migration versions (e.g., V1.0.0__description.sql).
	re := regexp.MustCompile(`V(\d+\.\d+\.\d+)__.*\.sql`)
	versions := make([]*sdk.ArtifactVersion, 0, len(files))
	for _, file := range files {
		filename := filepath.Base(file)
		if matches := re.FindStringSubmatch(filename); len(matches) > 1 {
			versions = append(versions, &sdk.ArtifactVersion{
				Kind:    sdk.ArtifactKindSQLMigration, // have to define this constant in sdk
				Version: matches[1],
				Name:    filename,
				URL:     file,
			})
		}
	}
	if len(versions) == 0 {
		return nil, fmt.Errorf("no valid migration versions found")
	}
	return versions, nil
}
