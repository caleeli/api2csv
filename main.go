package main

import (
	"github.com/caleeli/api2csv/internal/cli"
	"github.com/caleeli/api2csv/internal/storage/geo"
	"github.com/spf13/cobra"
)

func main() {
	// GeoRepo
	repo := geo.NewGeoRepo()
	// CsvRepo
	csvRepo := cli.NewCsvRepo()
	// cli
	rootCmd := &cobra.Command{Use: "api2csv-cli"}
	rootCmd.AddCommand(cli.InitializeCommand(repo, csvRepo))
	rootCmd.Execute()
}
