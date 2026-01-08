package main

import (
	"adventofcode/utils"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day08")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	lines := GetInput()

	fmt.Println("Day 08, Part 1 Answer: ", partOne(lines, 1000))
	fmt.Println("Day 08, Part 2 Answer: ", partTwo(lines))
}

// Point3D representerar en punkt i 3D-rymd (x, y, z)
type Point3D []float64

// JunctionBox representerar en elektrisk kopplingsbox med position i 3D
type JunctionBox struct {
	OriginalLine string  // Ursprungliga raden från input
	Index        int     // Index i listan (0, 1, 2, ...)
	Position     Point3D // 3D-koordinater
}

// DistanceTo beräknar det euklidiska avståndet mellan två punkter i 3D
// Formeln är: sqrt((x1-x2)² + (y1-y2)² + (z1-z2)²)
func (p Point3D) DistanceTo(other Point3D) float64 {
	sumOfSquares := 0.0

	for i := range p {
		diff := p[i] - other[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares)
}

// Connection representerar en möjlig koppling mellan två junction boxes
type Connection struct {
	BoxA     int     // Index för första boxen
	BoxB     int     // Index för andra boxen
	Distance float64 // Avståndet mellan dem
}

func partOne(lines []string, numberOfConnectionsToMake int) int {
	// Steg 1: Parsa input till junction boxes
	junctionBoxes := parseJunctionBoxes(lines)

	// Steg 2: Beräkna alla möjliga kopplingar mellan alla par av boxes
	// Om vi har 1000 boxes finns det 1000*999/2 = ~500,000 möjliga kopplingar
	allPossibleConnections := calculateAllConnections(junctionBoxes)

	// Steg 3: Sortera kopplingarna efter avstånd (kortast först)
	slices.SortFunc(allPossibleConnections, func(a, b Connection) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	// Steg 4: Använd Union-Find för att spåra vilka boxes som tillhör samma krets
	//
	// Union-Find är en datastruktur som effektivt håller koll på grupper/mängder.
	// - find(x): Hitta vilken grupp box x tillhör (returnerar "root" för gruppen)
	// - union(x, y): Slå ihop grupperna som x och y tillhör till en grupp
	//
	// Exempel: Om box 1 och 3 är i samma krets, och vi gör union(1, 5),
	// då blir box 1, 3 och 5 alla i samma krets.

	circuitTracker := newUnionFind(len(junctionBoxes))

	// Steg 5: Gör de N kortaste kopplingarna
	// (Uppgiften säger 1000 kopplingar för riktiga input)
	for i := 0; i < numberOfConnectionsToMake && i < len(allPossibleConnections); i++ {
		connection := allPossibleConnections[i]
		// Koppla ihop de två boxarna - de blir nu del av samma krets
		circuitTracker.union(connection.BoxA, connection.BoxB)
	}

	// Steg 6: Räkna hur stora de olika kretsarna är
	circuitSizes := circuitTracker.getGroupSizes()

	// Steg 7: Sortera storlekarna (störst först)
	slices.Sort(circuitSizes)
	slices.Reverse(circuitSizes)

	// Steg 8: Multiplicera de tre största kretsarnas storlekar
	result := 1
	for i := 0; i < 3 && i < len(circuitSizes); i++ {
		result *= circuitSizes[i]
	}

	return result
}

// calculateAllConnections beräknar avståndet mellan alla par av junction boxes
func calculateAllConnections(boxes []JunctionBox) []Connection {
	connections := []Connection{}

	// Loopa genom alla unika par (i, j) där i < j
	// Detta undviker dubbletter (vi vill inte ha både (1,2) och (2,1))
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			distance := boxes[i].Position.DistanceTo(boxes[j].Position)
			connections = append(connections, Connection{
				BoxA:     i,
				BoxB:     j,
				Distance: distance,
			})
		}
	}

	return connections
}

// UnionFind är en datastruktur för att spåra vilka element som tillhör samma grupp
type UnionFind struct {
	// parent[i] = föräldern till element i
	// Om parent[i] == i, då är i "root" (ledaren) för sin grupp
	parent []int

	// rank[i] = ungefärlig djup av trädet under i
	// Används för att hålla träden balanserade (optimering)
	rank []int
}

// newUnionFind skapar en ny Union-Find struktur där varje element startar i sin egen grupp
func newUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)

	// I början är varje element sin egen förälder (sin egen grupp)
	for i := range parent {
		parent[i] = i
	}

	return &UnionFind{parent: parent, rank: rank}
}

// find hittar "root" (ledaren) för gruppen som element x tillhör
// Använder "path compression" - alla element längs vägen pekar direkt på root efteråt
func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		// Rekursivt hitta root och uppdatera parent för snabbare framtida lookups
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

// union slår ihop grupperna som x och y tillhör
func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	// Om de redan är i samma grupp, gör inget
	if rootX == rootY {
		return
	}

	// Slå ihop grupperna - häng det mindre trädet under det större
	// Detta håller träden balanserade för snabbare find()
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

// getGroupSizes returnerar en lista med storleken på varje grupp
func (uf *UnionFind) getGroupSizes() []int {
	// Räkna hur många element som har varje root
	sizeByRoot := make(map[int]int)

	for i := range uf.parent {
		root := uf.find(i)
		sizeByRoot[root]++
	}

	// Konvertera map till lista
	sizes := []int{}
	for _, size := range sizeByRoot {
		sizes = append(sizes, size)
	}

	return sizes
}

// parseJunctionBoxes konverterar input-rader till JunctionBox-strukturer
func parseJunctionBoxes(lines []string) []JunctionBox {
	boxes := []JunctionBox{}

	for i, line := range lines {
		// Splitta "162,817,812" till ["162", "817", "812"]
		coordinates := strings.Split(line, ",")

		// Konvertera strängar till float64
		position := Point3D{}
		for _, coord := range coordinates {
			number, _ := strconv.ParseFloat(coord, 64)
			position = append(position, number)
		}

		boxes = append(boxes, JunctionBox{
			OriginalLine: line,
			Index:        i,
			Position:     position,
		})
	}

	return boxes
}

func partTwo(lines []string) int {
	total := len(lines)
	return total
}
