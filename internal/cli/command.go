package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// CobraFn function definition for Cobra commands
type CobraFn func(cmd *cobra.Command, args []string)

func InitializeCommand(repo GeoRepo, csvRepo CsvRepo) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "location",
		Short: "Get your current location",
		Run:   runGet(repo, csvRepo),
	}
	cmd.Flags().StringP("output", "o", "", "Output file")
	return cmd
}

func runGet(repo GeoRepo, csvRepo CsvRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		location, _ := repo.Locate()
		fmt.Printf("Your current location right now is: %s, %s\n", location.City, location.Country)
		if output != "" {
			file, err := os.Create(output)
			if err != nil {
				log.Fatalln("error creating file,", err)
			}
			csvRepo.SaveCSV(location, file)
		}
	}
}
