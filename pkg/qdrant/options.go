package qdrant

// WithReplicaCount ...
func (e Env) WithReplicaCount(replicas int) Env {
	e.ReplicaCount = int64(replicas)
	return e
}

// WithImageTag ...
func (e Env) WithImageTag(tag string) Env {
	e.ImageTag = tag
	return e
}

// WithResources ...
func (e Env) WithResources(cpu, memory string) Env {
	e.Resources = Resources{
		CPU:    cpu,
		Memory: memory,
	}
	return e
}
