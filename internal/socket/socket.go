package socket

import (
	"os"
	"path"
)

const (
	_socketPath = "workspace/server.sock"
	_envPath    = "SOCKET_PATH"
)

func MustGetPath() string {
	base := os.Getenv(_envPath)
	if len(base) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		base = cwd
	}

	return path.Join(base, _socketPath)
}

func MustEnsureSocket(socketPath string) {
	_, err := os.Stat(socketPath)
	if !os.IsNotExist(err) {
		os.Remove(socketPath)
		return
	}

	return
}
