package passcl

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

type OsClient struct{}

func NewOsClient() *OsClient {
	return &OsClient{}
}

func (c *OsClient) Init(ctx context.Context, subFolder *string, key string) error {
	var subFolderFlag string

	if subFolder != nil {
		subFolderFlag = fmt.Sprintf("-p %s", *subFolder)
	}

	return run(ctx, "pass", "init", subFolderFlag, key)
}

func (c *OsClient) Show(ctx context.Context, passName string) (string, error) {
	return out(ctx, "pass", "show", passName)
}

func (c *OsClient) ShowOtp(ctx context.Context, passName string) (string, error) {
	return out(ctx, "pass", "otp", "show", passName)
}

func (c *OsClient) Generate(ctx context.Context, passName string, options GenerateOptions) error {
	length := strconv.Itoa(options.Length)

	var noSymbolsFlag string
	if options.NoSymbols {
		noSymbolsFlag = "-n"
	}

	var forceFlag string
	if options.Force {
		forceFlag = "-f"
	}

	return run(ctx, "pass", "generate", noSymbolsFlag, forceFlag, passName, length)
}

func (c *OsClient) Insert(ctx context.Context, passName string, data []string, options InsertOptions) error {
	var forceFlag string
	if options.Force {
		forceFlag = "-f"
	}

	var multiLineFlag string
	if options.MultiLine {
		multiLineFlag = "-m"
	} else {
		data = []string{data[0], data[0]}
	}

	return prompt(ctx, "pass", data, "insert", multiLineFlag, forceFlag, passName)
}

func (c *OsClient) InsertOtp(ctx context.Context, passName string, uri string, options InsertOtpOptions) error {
	var forceFlag string
	if options.Force {
		forceFlag = "-f"
	}

	return prompt(ctx, "pass", []string{uri, uri}, "otp", "insert", forceFlag, passName)
}

func run(ctx context.Context, bin string, args ...string) error {
	cmd := buildCmd(ctx, bin, args...)
	return osErr(cmd.Run())
}

func out(ctx context.Context, bin string, args ...string) (string, error) {
	cmd := buildCmd(ctx, bin, args...)

	rawOut, err := cmd.Output()
	if err != nil {
		return "", osErr(err)
	}

	return string(rawOut), nil
}

func prompt(ctx context.Context, bin string, dataToInput []string, args ...string) error {
	cmd := buildCmd(ctx, bin, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return osErr(err)
	}

	if err = cmd.Start(); err != nil {
		return osErr(err)
	}

	for _, input := range dataToInput {
		byteInput := []byte(fmt.Sprintf("%s\n", input))

		_, err = stdin.Write(byteInput)
		if err != nil {
			return osErr(err)
		}
	}

	_ = stdin.Close()

	if err = cmd.Wait(); err != nil {
		return osErr(err)
	}

	return nil
}

func buildCmd(ctx context.Context, bin string, args ...string) *exec.Cmd {
	cmdArgs := buildCmdArgsFlags(args...)
	return exec.CommandContext(ctx, bin, cmdArgs...)
}

func buildCmdArgsFlags(args ...string) []string {
	res := make([]string, 0, len(args))

	for _, arg := range args {
		if arg == "" {
			continue
		}

		res = append(res, arg)
	}

	return res
}

func osErr(err error) error {
	var exitErr *exec.ExitError
	ok := errors.As(err, &exitErr)
	if ok {
		outErr := exitErr.Stderr
		return errors.New(string(outErr))
	} else {
		return err
	}
}
