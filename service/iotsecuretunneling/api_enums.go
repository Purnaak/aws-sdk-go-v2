// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsecuretunneling

type ConnectionStatus string

// Enum values for ConnectionStatus
const (
	ConnectionStatusConnected    ConnectionStatus = "CONNECTED"
	ConnectionStatusDisconnected ConnectionStatus = "DISCONNECTED"
)

func (enum ConnectionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TunnelStatus string

// Enum values for TunnelStatus
const (
	TunnelStatusOpen   TunnelStatus = "OPEN"
	TunnelStatusClosed TunnelStatus = "CLOSED"
)

func (enum TunnelStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TunnelStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}