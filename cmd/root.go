package cmd

import (
	"bufio"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	missingInputFileError = errors.New("specified input file does not exist")
	readInputFileError    = errors.New("failed to read the input file")
)

var log = logrus.New()

type world struct {
	width, height int
}

type command struct {
	direction string
}

type robot struct {
	commands []command
}

type universe struct {
	world  *world
	robots []robot
}

var rootCmd = cobra.Command{
	Short: "Martian Robots app",
	Long:  `This CLI app determines each sequence of robot positions and reports the final position of the robot.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ := cmd.Flags().GetString("input")
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			log.Fatal(missingInputFileError)
		}

		file, err := os.Open(inputPath)
		if err != nil {
			log.Fatal(readInputFileError)
		}
		defer file.Close()

		var (
			scanner  = bufio.NewScanner(file)
			universe = &universe{}
		)
		var line int
		for scanner.Scan() {
			if line == 0 {

			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	rootCmd.Flags().Bool("debug", false, "Run in debug mode")
	rootCmd.Flags().String("input", "", "Specify an input path")
	_ = rootCmd.MarkFlagRequired("input")
}
