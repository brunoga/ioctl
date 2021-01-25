package ioctl

import (
	"fmt"

	"github.com/brunoga/ioctl/uapi"
	"golang.org/x/sys/unix"
)

type Command interface {
	Direction() uintptr
	Type() uintptr
	Number() uintptr
	Size() uintptr

	Do(fd, arg uintptr) error
}

type commandImpl uintptr

func NewReadCommand(typ, nr, size uintptr) Command {
	return commandImpl(uapi.Ior(typ, nr, size))
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

func (ci commandImpl) Do(fd, arg uintptr) error {
	return ioctl(fd, uintptr(ci), arg)
}

func ioctl(fd, req, arg uintptr) (err error) {
	_, _, e1 := unix.Syscall(unix.SYS_IOCTL, fd, req, arg)
	if e1 != 0 {
		return fmt.Errorf("error %d", e1)
	}

	return nil
}

