package qdrant

import "errors"

var (
	ErrFailedToMarshalValues = errors.New("failed to marshal values to YAML")
	ErrFailedToWriteFile     = errors.New("failed to write values file")
)
