package qdrant

// DefaultEnv ...
func DefaultEnv(replicas int) Env {
	return Env{
		ReplicaCount: int64(replicas),
		ImageTag:     "latest",
		Resources: Resources{
			CPU:    "500m",
			Memory: "512Mi",
		},
	}
}
