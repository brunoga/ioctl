// +build linux

// Package ioctl implements the Linux userland (UAPI) ioctl interface for
// generating and parsing ioctl command definitions.
package ioctl

const (
	iocNrBits   uintptr = 8
	iocTypeBits uintptr = 8
	iocSizeBits uintptr = 14
	iocDirBits  uintptr = 2

	iocNrMask   uintptr = ((1 << iocNrBits) - 1)
	iocTypeMask uintptr = ((1 << iocTypeBits) - 1)
	iocSizeMask uintptr = ((1 << iocSizeBits) - 1)
	iocDirMask  uintptr = ((1 << iocDirBits) - 1)

	iocNrShift   uintptr = 0
	iocTypeShift uintptr = iocNrShift + iocNrBits
	iocSizeShift uintptr = iocTypeShift + iocTypeBits
	iocDirShift  uintptr = iocSizeShift + iocSizeBits

	iocNone  uintptr = 0
	iocWrite uintptr = 1
	iocRead  uintptr = 2
)

func ioc(dir, typ, nr, size uintptr) uintptr {
	return ((dir << iocDirShift) |
		(typ << iocTypeShift) |
		(nr << iocNrShift) |
		(size << iocSizeShift))
}

// Io creates an ioctl command definition that does not read data from or write
// data to the kernel.
func Io(typ, nr uintptr) uintptr {
	return ioc(iocNone, typ, nr, 0)
}

// Ior creates an ioctl command definition that reads data from the kernel.
func Ior(typ, nr, size uintptr) uintptr {
	return ioc(iocRead, typ, nr, size)
}

// Iow creates an ioctl command definition that writes data to the kernel.
func Iow(typ, nr, size uintptr) uintptr {
	return ioc(iocWrite, typ, nr, size)
}

// Iowr creates an ioctl command definition that reads data from and writes
// data to the kernel.
func Iowr(typ, nr, size uintptr) uintptr {
	return ioc((iocRead | iocWrite), typ, nr, size)
}

// IocDir returns the direction associated with the given command definition.
func IocDir(nr uintptr) uintptr {
	return ((nr >> iocDirShift) & iocDirMask)
}

// IocType returns the type associated with the given command definition.
func IocType(nr uintptr) uintptr {
	return ((nr >> iocTypeShift) & iocTypeMask)
}

// IocNr returns the command number associated with the given command
// definition.
func IocNr(nr uintptr) uintptr {
	return ((nr >> iocNrShift) & iocNrMask)
}

// IocSize returns the size associated with the given command definition.
func IocSize(nr uintptr) uintptr {
	return ((nr >> iocSizeShift) & iocSizeMask)
}
