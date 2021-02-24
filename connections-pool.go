package pool


AbstractPool interface {
	get() (fd File)
	free(fd Free)
}

// We want thread safety
VerticalConnectionsPool struct {
	connections chan FD
}


func NewVerticalConnectionsPool(size int) AbstractPool {
	pool := &VerticalConnectionsPool {
		connections: make chan FD
	}
	pool.openConnections()

	return pool
}

(vcp *VerticalConnectionsPool) get() (fd File) {
	fd = <- vcp.connections
	return
}

(vcp *VerticalConnectionsPool) free(fd Free) {
	vcp.connections <- fd
}

