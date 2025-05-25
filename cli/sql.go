package cli

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type options struct {
	// databaseURL string 
	// user        string 
	// password   string 
}

type Option func(*options)

type SQLClient struct {
	execPath   string
	dir        string
	configPath string
	options    options
}

func NewSQLClient(execPath, dir, configPath string, opts ...Option) *SQLClient {
	opt := options{}
	for _, o := range opts {
		o(&opt)
	}
	return &SQLClient{
		execPath:   execPath,
		dir:        dir,
		configPath: configPath,
		options:    opt,
	}
}

func (s *SQLClient) Version(ctx context.Context) (string, error) {
	args := []string{"--version"}
	cmd := exec.CommandContext(ctx, s.execPath, args...)
	cmd.Dir = s.dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return strings.TrimSpace(string(out)), nil
}

func (s *SQLClient) Apply(ctx context.Context, w io.Writer) error {
	args := []string{
		"-f", s.configPath, // Execute SQL script file
	}
	cmd := exec.CommandContext(ctx, s.execPath, args...)
	cmd.Dir = s.dir
	cmd.Stdout = w
	cmd.Stderr = w
	fmt.Fprintf(w, "execute: '%s %s'\n", s.execPath, strings.Join(args, " "))
	return cmd.Run()
}

func (s *SQLClient) Diff(ctx context.Context, w io.Writer) error {
	args := []string{
		"--version", // will be replaced with actual diff command if available (pg_diff)
	}
	cmd := exec.CommandContext(ctx, s.execPath, args...)
	cmd.Dir = s.dir
	cmd.Stdout = w
	cmd.Stderr = w
	fmt.Fprintf(w, "execute: '%s %s'\n", s.execPath, strings.Join(args, " "))
	return cmd.Run()
}