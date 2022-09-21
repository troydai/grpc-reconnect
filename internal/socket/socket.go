package socket

import (
	"os"
	"path"
)

const _socketPath = "workspace/server.sock"

func MustGetPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path.Join(cwd, _socketPath)
}

func MustEnsureSocket(socketPath string) {
	_, err := os.Stat(socketPath)
	if !os.IsNotExist(err) {
		os.Remove(socketPath)
		return
	}

	return
}
