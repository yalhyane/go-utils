package net_utils

import "net"

// GetFreePort resolves a free port on the local machine and returns it, or an error if one occurs.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func(l *net.TCPListener) {
		_ = l.Close()
	}(l)
	return l.Addr().(*net.TCPAddr).Port, nil
}
