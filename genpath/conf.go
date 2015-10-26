package genpath
import (
	"errors"
	"fmt"
)

const (
	DEFAULT_POPULATION_SIZE = 10
	DEFAULT_CROSSOVERS_COUNT = 1
	DEFAULT_CROSSOVER_PROB = 1.0
	DEFAULT_MUTATION_PROB = 0.1
	DEFAULT_SELECTION_COUNT = 2
	DEFAULT_CROSSOVER_SEGMENT_SPLITS_COUNT = 1
)

type FitnessFunc func(*Genome) FitnessT
type GraphT [][]float64

type GenPathConf struct {
	Graph GraphT
	StartNode int
	EndNode int

	PopulationSize int
	CrossoversCount int
	CrossoverProb float64
	CrossoverSegmentSplitsCount int
	SelectionCount int
	RemoveDuplicates bool

	MutationProb float64

	fitnessFunc FitnessFunc
}

func (self *GenPathConf) check() error {
	if self.Graph == nil {
		return errors.New("Graph not supplied")
	}

	if len(self.Graph) == 0 || len(self.Graph) != len(self.Graph[0]) {
		return errors.New("Graph has incorrect structure")
	}

	if self.PopulationSize == 0 {
		self.PopulationSize = DEFAULT_POPULATION_SIZE
	}

	if (self.StartNode < 0 || self.StartNode >= len(self.Graph)) {
		return errors.New(fmt.Sprintf("Start node has to be >= 0 and < %d. Got: %d", len(self.Graph), self.StartNode))
	}

	if (self.EndNode < 0) {
		return errors.New(fmt.Sprintf("End node has to be >= 0 and < %d. Got: %d", len(self.Graph), self.EndNode))
	}

	if self.CrossoversCount <= 0 {
		self.CrossoversCount = DEFAULT_CROSSOVERS_COUNT
	}

	if self.CrossoverProb < 0.0 {
		self.CrossoverProb = DEFAULT_CROSSOVER_PROB
	}

	if self.CrossoverSegmentSplitsCount <= 0.0 {
		self.CrossoverProb = DEFAULT_CROSSOVER_SEGMENT_SPLITS_COUNT
	}

	if self.SelectionCount <= 0 {
		self.SelectionCount = DEFAULT_SELECTION_COUNT
	}

	if self.MutationProb < 0.0 {
		self.MutationProb = DEFAULT_MUTATION_PROB
	}


	if self.SelectionCount > self.PopulationSize {
		self.SelectionCount = self.PopulationSize
	}

	if self.fitnessFunc == nil {
		panic("No fitness function in conf")
	}

	return nil
}

func (self *GenPathConf) dist(node1, node2 int) float64 {
	return self.Graph[node1][node2]
}