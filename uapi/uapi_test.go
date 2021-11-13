//go:build linux

// Package uapi implements the Linux userland ioctl interface for generating
// and parsing ioctl command definitions.
package uapi

import "testing"

func Test_ioc(t *testing.T) {
	type args struct {
		dir  uintptr
		typ  uintptr
		nr   uintptr
		size uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"WithZeroes", args{0, 0, 0, 0}, 0},
		{"WithData", args{1, 2, 3, 4}, 1074004483},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ioc(tt.args.dir, tt.args.typ, tt.args.nr, tt.args.size); got != tt.want {
				t.Errorf("ioc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIo(t *testing.T) {
	type args struct {
		typ uintptr
		nr  uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"WithZeroes", args{0, 0}, 0},
		{"WithData", args{1, 2}, 258},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Io(tt.args.typ, tt.args.nr); got != tt.want {
				t.Errorf("Io() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIor(t *testing.T) {
	type args struct {
		typ  uintptr
		nr   uintptr
		size uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"WithZeroes", args{0, 0, 0}, 2147483648},
		{"WithData", args{1, 2, 3}, 2147680514},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ior(tt.args.typ, tt.args.nr, tt.args.size); got != tt.want {
				t.Errorf("Ior() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIow(t *testing.T) {
	type args struct {
		typ  uintptr
		nr   uintptr
		size uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"WithZeroes", args{0, 0, 0}, 1073741824},
		{"WithData", args{1, 2, 3}, 1073938690},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iow(tt.args.typ, tt.args.nr, tt.args.size); got != tt.want {
				t.Errorf("Iow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIowr(t *testing.T) {
	type args struct {
		typ  uintptr
		nr   uintptr
		size uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"WithZeroes", args{0, 0, 0}, 3221225472},
		{"WithData", args{1, 2, 3}, 3221422338},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iowr(tt.args.typ, tt.args.nr, tt.args.size); got != tt.want {
				t.Errorf("Iowr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIocDir(t *testing.T) {
	type args struct {
		nr uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"None", args{Io(0, 0)}, 0},
		{"Write", args{Iow(0, 0, 0)}, 1},
		{"Read", args{Ior(0, 0, 0)}, 2},
		{"WriteRead", args{Iowr(0, 0, 0)}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IocDir(tt.args.nr); got != tt.want {
				t.Errorf("IocDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIocType(t *testing.T) {
	type args struct {
		nr uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"Zero", args{Io(0, 1)}, 0},
		{"One", args{Io(1, 1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IocType(tt.args.nr); got != tt.want {
				t.Errorf("IocType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIocNr(t *testing.T) {
	type args struct {
		nr uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"Zero", args{Io(1, 0)}, 0},
		{"One", args{Io(1, 1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IocNr(tt.args.nr); got != tt.want {
				t.Errorf("IocNr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIocSize(t *testing.T) {
	type args struct {
		nr uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"Zero", args{Ior(1, 2, 0)}, 0},
		{"One", args{Ior(1, 2, 1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IocSize(tt.args.nr); got != tt.want {
				t.Errorf("IocSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
