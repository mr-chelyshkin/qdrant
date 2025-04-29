package qdrant

// Env ...
type Env struct {
	ReplicaCount int64
	ImageTag     string
	Resources    Resources
}

// Resources ...
type Resources struct {
	CPU    string
	Memory string
}

// ToMap ...
func (e Env) ToMap() map[string]any {
	return map[string]any{
		"replicaCount": e.ReplicaCount,
		"image": map[string]any{
			"tag": e.ImageTag,
		},
		"resources": map[string]any{
			"limits": map[string]any{
				"cpu":    e.Resources.CPU,
				"memory": e.Resources.Memory,
			},
			"requests": map[string]any{
				"cpu":    e.Resources.CPU,
				"memory": e.Resources.Memory,
			},
		},
	}
}
