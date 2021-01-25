package ioctl

import (
	"fmt"
	"reflect"

	"github.com/brunoga/ioctl/uapi"
	"golang.org/x/sys/unix"
)

// Command is a higher level abstraction of a command definition that is also
// able to execute them.
type Command interface {
	// Direction returns what the kernel does to the arg parameter for Do():
	// Nothing, read, write or read and write.
	Direction() uintptr

	// Type identifies a specific subsystem or driver.
	Type() uintptr

	// Number identifies the specific command and is unique for a
	// given Type().
	Number() uintptr

	// Size of the underlying type of the arg Do() parameter.
	Size() uintptr

	// Do executes the command on the given fd passing the given arg to the
	// underlying ioctl call. The arg parameter must be a pointer type.
	Do(fd uintptr, arg interface{}) error
}

type commandImpl uintptr

// NewNoneCommand returns a Command that does not have any arguments.
func NewNoneCommand(typ, nr uintptr) Command {
	return commandImpl(uapi.Io(typ, nr))
}

// NewReadCommand returns a Command that reads data from the kernel.
func NewReadCommand(typ, nr, size uintptr) Command {
	return commandImpl(uapi.Ior(typ, nr, size))
}

// NewWriteCommand returns a Command that writes (sends) data to the kernel.
func NewWriteCommand(typ, nr, size uintptr) Command {
	return commandImpl(uapi.Iow(typ, nr, size))
}

// NewReadWriteCommand returns a Command that both reads and writes (sends)
// data to the kernel.
func NewReadWriteCommand(typ, nr, size uintptr) Command {
	return commandImpl(uapi.Iowr(typ, nr, size))
}

func (ci commandImpl) Direction() uintptr {
	return uapi.IocDir(uintptr(ci))
}

func (ci commandImpl) Type() uintptr {
	return uapi.IocType(uintptr(ci))
}

func (ci commandImpl) Number() uintptr {
	return uapi.IocNr(uintptr(ci))
}

func (ci commandImpl) Size() uintptr {
	return uapi.IocSize(uintptr(ci))
}

func (ci commandImpl) Do(fd uintptr, arg interface{}) error {
	argVal := reflect.ValueOf(arg)

	if argVal.Kind() != reflect.Ptr {
		return fmt.Errorf("pointer type is required")
	}

	return ioctl(fd, uintptr(ci), argVal.Pointer())
}

func ioctl(fd, req, arg uintptr) (err error) {
	_, _, e1 := unix.Syscall(unix.SYS_IOCTL, fd, req, arg)
	if e1 != 0 {
		return fmt.Errorf("error %d", e1)
	}

	return nil
}
