package main

// Strings value are immutable
// Strings can be compared with comparision operators like == and <
// The comparision is done byte by byte.
// Immutability means that it is safe for two copies of a string to share the same underlying memory,

// A string is a sequence of bytes, and each character within the string may consist of multiple bytes
// A rune, on the other hand, represents a single Unicode character regardless of its byte representation.
type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast
	// supports broadcast access capability
	FlagLoopback
	// is a loopback interface
	FlagPointToPoint
	// belongs to a point-to-point link
	FlagMulticast

// supports multicast access capability
)

func main() {
	const GoUsage = `Go is a tool for managing Go source code.
	Usage:
	go command [arguments]
	...`
}
