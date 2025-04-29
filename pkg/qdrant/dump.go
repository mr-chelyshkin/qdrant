package qdrant

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// WriteValuesToFile ...
func WriteValuesToFile(values map[string]any, path string) error {
	data, err := yaml.Marshal(values)
	if err != nil {
		return errors.Join(ErrFailedToMarshalValues, err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return errors.Join(ErrFailedToWriteFile, err)
	}

	return nil
}
