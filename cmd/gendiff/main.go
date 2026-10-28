// Команда gendiff сравнивает два конфигурационных файла.
package main

import (
	"context"
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		ArgsUsage: "<filepath1> <filepath2>",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Value: "stylish", Usage: "output format"},
		},
		Action: run,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(_ context.Context, cmd *cli.Command) error {
	if cmd.NArg() != 2 {
		return fmt.Errorf("expected two file paths, got %d", cmd.NArg())
	}

	result, err := code.GenDiff(cmd.Args().Get(0), cmd.Args().Get(1), cmd.String("format"))
	if err != nil {
		return fmt.Errorf("gendiff: %w", err)
	}

	fmt.Println(result)

	return nil
}
