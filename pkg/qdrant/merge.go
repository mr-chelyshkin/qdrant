package qdrant

import "maps"

// MergeValues ...
func MergeValues(base, override map[string]any) map[string]any {
	out := make(map[string]any, len(base))

	maps.Copy(out, base)
	maps.Copy(out, override)

	return out
}
