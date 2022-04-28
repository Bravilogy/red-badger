package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/bravilogy/robots/domain"
)

var (
	missingInputFileError = errors.New("specified input file does not exist")
	readInputFileError    = errors.New("failed to read the input file")
)

var log = logrus.New()

var rootCmd = cobra.Command{
	Short: "Martian Robots app",
	Long:  `This CLI app determines each sequence of robot positions and reports the final position of the robot.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ := cmd.Flags().GetString("input")
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			log.Fatal(missingInputFileError)
		}

		rawInput, err := os.ReadFile(inputPath)
		if err != nil {
			log.Fatal(readInputFileError)
		}

		universe, err := domain.NewUniverseFromString(string(rawInput))
		if err != nil {
			log.Fatal(err)
		}

		// run through each robot commands and apply them to the world
		for _, robot := range universe.Robots {
			for _, command := range robot.Commands {
				if robot.IsLost() {
					continue
				}

				command.Run(robot, universe.World)
			}

			fmt.Println(robot.Report())
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
