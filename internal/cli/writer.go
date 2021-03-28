package cli

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type CsvRepo interface {
	SaveCSV(location *Location, output *os.File)
}

type csvRepo struct {
	separator rune
}

func NewCsvRepo() CsvRepo {
	return &csvRepo{','}
}

func (c *csvRepo) SaveCSV(location *Location, output *os.File) {
	writer := csv.NewWriter(output)
	writer.Comma = c.separator

	row := []string{
		location.CountryCode,
		location.Country,
		location.Region,
		location.City,
		location.ZipCode,
		strconv.FormatFloat(location.Latitude, 'f', 5, 64),
		strconv.FormatFloat(location.Longitude, 'f', 5, 64),
	}
	err := writer.Write(row)
	if err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
