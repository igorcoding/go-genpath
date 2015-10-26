package main
import (
	"github.com/igorcoding/go-genpath/genpath"
	"github.com/kr/pretty"
	"fmt"
"math"
	"math/rand"
)


func main() {
	rand.Seed(1234);
	inf := math.Inf(0)
	graph := [][]float64 {
		{0, 2, 1, inf, 10},
		{2, 0, 4, 3, 1},
		{1, 4, 0, 1, inf},
		{inf, 3, 1, 0, 1},
		{10, 1, inf, 1, 0},
	};
	conf := &genpath.GenPathConf{
		PopulationSize: 15,
		MutationProb: 1.0,
		Graph: graph,
		StartNode: 0,
		EndNode: 4,
	}
	gPath, err := genpath.NewGenPath(conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < 100; i++ {
		pretty.Println(i)
		gPath.Step()
	}

	pretty.Println(gPath)
}