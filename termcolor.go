package termcolor

//go:build darwin || dragonfly || freebsd || netbsd || openbsd || linux || (!windows && !js)
// +build darwin dragonfly freebsd netbsd openbsd linux !windows,!js

import (
	"syscall"
	"unsafe"
)

//go:build darwin || dragonfly || freebsd || netbsd || openbsd
// +build darwin dragonfly freebsd netbsd openbsd
const ioctlReadTermios = syscall.TIOCGETA // for bsd based

const ioctlReadTermios = 0x5401 // syscall.TCGETS // for linux
const fadviseDontneed = 4

/* defined in linux-4.14/include/uapi/linux/fadvise.h
 * #define POSIX_FADV_DONTNEED 4
 */

//go:build linux
// +build linux

func fadvise(fd int, offset int64, length int64, advice int) (err error) {
	return unix.Fadvise(fd, offset, length, advice)
}

func TryToDropFilePageCache(fd int, offset int64, length int64) {
	fadvise(fd, offset, length, fadviseDontneed)
}

func compareFileCreatedTime(a, b os.FileInfo) bool {
	stati := a.Sys().(*syscall.Stat_t)
	statj := b.Sys().(*syscall.Stat_t)
	ctimei := time.Unix(int64(stati.Ctim.Sec), int64(stati.Ctim.Nsec))
	ctimej := time.Unix(int64(statj.Ctim.Sec), int64(statj.Ctim.Nsec))
	return ctimei.After(ctimej)
}

// p.color = runtime.GOOS != "windows" && IsTerminal(int(os.Stdout.Fd()))
func IsTerminal(fd int) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
