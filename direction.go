package ioctl

type Direction uintptr

const (
	// DirectionNone means that the kernel does not read or write anything.
	DirectionNone Direction = 0

	// DirectionRead means that the kernel reads data from the command arg
	// parameter.
	DirectionRead Direction = 1

	// DirectionWrite means that the kernel writes data to the command arg
	// parameter.
	DirectionWrite Direction = 2

	// DirectionReadWrite means that the kernel reads and writes data to the
	// command arg parameter.
	DirectionReadWrite Direction = 3
)
