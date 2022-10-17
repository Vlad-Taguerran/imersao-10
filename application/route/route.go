package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID       string
	ClientId string
	Position []Position
}

type Position struct {
	Lat  float64
	Long float64
}

type PartialRoutePosition struct {
	ID       string    `json:"RouteId"`
	ClientId string    `json:"ClientId"`
	Position []float64 `json:"Position"`
	Finished bool      `json:"Finished"`
}

// LoadPositions Metodo que carrega a posição apartir de arquivos txt
func (route *Route) LoadPositions() error {

	if route.ID == "" {
		return errors.New("ID vazio")
	}
	// leitura de arquivos txt
	file, err := os.Open("./")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}

		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}

		route.Position = append(route.Position, Position{
			Lat:  lat,
			Long: long,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Position)

	for k, v := range r.Position {

		route.ID = r.ID
		route.ClientId = r.ClientId
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false

		if total-1 == k {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}
