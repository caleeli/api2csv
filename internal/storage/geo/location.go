package geo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/caleeli/api2csv/internal/cli"
)

type geoRepo struct {
	url string
}

// NewGeoRepo creates a new google TransRepo
func NewGeoRepo() cli.GeoRepo {
	return &geoRepo{"https://freegeoip.app/json/"}
}

func (repo *geoRepo) Locate() (translation *cli.Location, err error) {
	response, err := http.Get(repo.url)
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &translation)
	if err != nil {
		return nil, err
	}
	return
}
