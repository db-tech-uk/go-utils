package osutils

import (
	"bytes"
	"context"
	"math"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// unixOS is the set of GOOS values matched by the "unix" build tag.
var unixOS = map[string]bool{
	"aix":       true,
	"android":   true,
	"darwin":    true,
	"dragonfly": true,
	"freebsd":   true,
	"hurd":      true,
	"illumos":   true,
	"ios":       true,
	"linux":     true,
	"netbsd":    true,
	"openbsd":   true,
	"solaris":   true,
}

func GetOS() string {
	return runtime.GOOS
}

func GetArch() string {
	return runtime.GOARCH
}

func OSIsUnix() bool {
	os := GetOS()

	if _, ok := unixOS[os]; ok {
		return true
	}

	return false
}

func RunCommand(command string, args ...string) (string, string, error) {
	return RunCommandContext(context.Background(), command, args...)
}

func RunCommandContext(ctx context.Context, command string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func RunCommandSilent(command string, args ...string) (string, string, error) {
	return RunCommandSilentContext(context.Background(), command, args...)
}

func RunCommandSilentContext(ctx context.Context, command string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, command, args...)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func RunShellCommand(shell string, command string) (string, string, error) {
	return RunShellCommandContext(context.Background(), shell, command)
}

func RunShellCommandContext(ctx context.Context, shell string, command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, shell, "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func RunShellCommandSilent(shell string, command string) (string, string, error) {
	return RunShellCommandSilentContext(context.Background(), shell, command)
}

func RunShellCommandSilentContext(ctx context.Context, shell string, command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, shell, "-c", command)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)

	return err == nil
}

func GetULimit() (int, error) {
	if !OSIsUnix() {
		// @TODO figure out what to do on non unix
		return 0, nil
	}

	var rLimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return 0, err
	}

	// Cur may be RLIM_INFINITY (max uint64); clamp rather than overflow.
	if rLimit.Cur > math.MaxInt {
		return math.MaxInt, nil
	}

	return int(rLimit.Cur), nil
}
