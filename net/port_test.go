package net_utils

import (
	"net"
	"testing"
)

func TestGetFreePort(t *testing.T) {
	// Get a free port.
	port, err := GetFreePort()
	if err != nil {
		t.Errorf("Failed to get a free port: %v", err)
	}

	// Verify that the port is actually free.
	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: port,
	})
	if err != nil {
		t.Errorf("Port %d is not free: %v", port, err)
	}
	defer func(l *net.TCPListener) {
		_ = l.Close()
	}(l)
}
