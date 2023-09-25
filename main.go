package main

import (
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/olekukonko/tablewriter"
)

// City representa una ciudad con nombre, coordenadas x e y.
type City struct {
	Name       string
	X, Y       float64
	connection []string
}

// Distance representa la distancia entre dos ciudades.
type Distance struct {
	To       string
	Distance float64
}

// Graph representa el grafo de ciudades y distancias.
type Graph struct {
	Cities      []City
	Connections map[string][]Distance
}

// NewGraph crea un nuevo grafo a partir de un arreglo de ciudades.
func NewGraph(cities []City) Graph {
	connections := make(map[string][]Distance)

	for _, city := range cities {
		for _, cityName := range city.connection {
			cityFind := findCity(cities, cityName)
			if cityFind != nil {
				connections[cityName] = append(connections[cityName], Distance{
					To:       city.Name,
					Distance: calculateDistance(city, *cityFind),
				})
			}
		}
	}

	return Graph{
		Cities:      cities,
		Connections: connections,
	}
}

// calculateDistance calcula la distancia entre dos ciudades.
func calculateDistance(city1, city2 City) float64 {
	return math.Sqrt(math.Pow(city1.X-city2.X, 2) + math.Pow(city1.Y-city2.Y, 2))
}

// Astar encuentra la ruta más cercana desde el inicio hasta el destino en el grafo.
func Astar(graph Graph, start, end string) ([]string, error) {
	// Verificamos si las ciudades de inicio y destino existen en el grafo.
	startCity, endCity := findCity(graph.Cities, start), findCity(graph.Cities, end)
	if startCity == nil || endCity == nil {
		return nil, fmt.Errorf("Ciudad de inicio o destino no encontrada")
	}

	openSet := make(map[string]bool)
	openSet[start] = true
	closedSet := make(map[string]bool)

	g := make(map[string]float64)
	for _, city := range graph.Cities {
		g[city.Name] = math.Inf(1)
	}
	g[start] = 0

	parent := make(map[string]string)

	for len(openSet) > 0 {

		current := findLowestF(openSet, g, end, &graph)

		if current == end {
			return buildPath(parent, end), nil
		}

		delete(openSet, current)
		closedSet[current] = true

		for _, neighbor := range graph.Connections[current] {

			if closedSet[neighbor.To] {
				continue
			}

			tentativeG := g[current] + neighbor.Distance

			if !openSet[neighbor.To] || tentativeG < g[neighbor.To] {
				parent[neighbor.To] = current
				g[neighbor.To] = tentativeG
				openSet[neighbor.To] = true
			}
		}
	}

	return nil, errors.New("No se encontró un camino válido")
}

// findLowestF encuentra la ciudad con el valor F más bajo en el conjunto abierto.
func findLowestF(openSet map[string]bool, g map[string]float64, end string, graph *Graph) string {
	lowestF := math.Inf(1)
	var lowestCity string

	for city := range openSet {
		f := g[city] + heuristic(city, end, graph)
		if f < lowestF {
			lowestF = f
			lowestCity = city
		}
	}

	return lowestCity
}

// findCity encuentra una ciudad en un arreglo de ciudades.
func findCity(cities []City, name string) *City {
	for _, city := range cities {
		if city.Name == name {
			return &city
		}
	}
	return nil
}

// buildPath construye la ruta desde el inicio hasta el destino.
func buildPath(parent map[string]string, current string) []string {
	var path []string
	for current != "" {
		path = append([]string{current}, path...)
		current = parent[current]
	}
	return path
}

// heuristic calcula la distancia entre dos ciudades.
func heuristic(city, end string, graph *Graph) float64 {
	startCity := findCity(graph.Cities, city)
	endCity := findCity(graph.Cities, end)
	if startCity != nil && endCity != nil {
		return math.Sqrt(math.Pow(startCity.X-endCity.X, 2) + math.Pow(startCity.Y-endCity.Y, 2))
	}
	return math.Inf(1)
}

// main es la función principal del programa.
func main() {
	graph := NewGraph([]City{
		{Name: "A", X: 0, Y: 0, connection: []string{"B", "C"}},
		{Name: "B", X: 1, Y: 1, connection: []string{"A", "D", "E"}},
		{Name: "C", X: 2, Y: 2, connection: []string{"A", "D", "E", "H"}},
		{Name: "D", X: 5, Y: 7, connection: []string{"C", "B", "F"}},
		{Name: "E", X: 3, Y: 3, connection: []string{"C", "B", "G"}},
		{Name: "F", X: 10, Y: 8, connection: []string{"D", "G"}},
		{Name: "G", X: 13, Y: 3, connection: []string{"F", "H", "E"}},
		{Name: "H", X: 13, Y: 13, connection: []string{"G", "C"}},
	})

	// Crear una nueva tabla.
	table := tablewriter.NewWriter(os.Stdout)

	// Definir los encabezados de la tabla.
	table.SetHeader([]string{"City", "X", "Y", "Connections"})

	// Agregar datos a la tabla.
	for _, city := range graph.Cities {
		connectionNames := ""
		if connections, ok := graph.Connections[city.Name]; ok {
			for _, conn := range connections {
				connectionNames += fmt.Sprintf("%s (%.2f), ", conn.To, conn.Distance)
			}
		}
		table.Append([]string{city.Name, fmt.Sprintf("%.2f", city.X), fmt.Sprintf("%.2f", city.Y), connectionNames})
	}

	// Renderizar la tabla.
	table.Render()

	// Encontrar la ruta más cercana desde "A" hasta "D".
	path, err := Astar(graph, "A", "F")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Ruta más cercana:", path)
	}
}
