package genpath
import (
	"sort"
	"math/rand"
	"github.com/kr/pretty"
"math"
)

type Population struct {
	conf *GenPathConf
	P Genomes
}

func NewPopulation(conf *GenPathConf) *Population {
	self := &Population{
		conf: conf,
		P: NewGenomes(conf, conf.PopulationSize),
	}

	for i := range(self.P) {
		self.P[i].EvalFitness()
	}
	self.sort()
	return self
}

func (self *Population) evaluate() {
	for k := 0; k < self.conf.CrossoversCount; k++ {
		if rand.Float64() < self.conf.CrossoverProb {
			g1 := self.P[rand.Int31n(int32(math.Min(float64(len(self.P)), float64(self.conf.SelectionCount))))]
			g2 := self.P[rand.Int31n(int32(math.Min(float64(len(self.P)), float64(self.conf.SelectionCount))))]

			child := g1.Crossover(g2)

			if rand.Float64() < self.conf.MutationProb {
				child.Mutate()
			}
			child.EvalFitness()

			self.P = append(self.P, child)
		}
	}

	self.sort()
	if (self.conf.RemoveDuplicates) {
		self.removeDuplicates()
	}
	if len(self.P) > self.conf.PopulationSize {
		self.P = self.P[:self.conf.PopulationSize]
	}
//	pretty.Println(self.P)
}

func (self *Population) sort() {
	sort.Sort(self.P)
}

func (self *Population) removeDuplicates() {
	p := make(Genomes, self.conf.PopulationSize)
	uniqueGenomes := make(map[string]bool)

	j := 0
	for i := range(self.P) {
		if j < self.conf.PopulationSize {
			s := self.P[i].ToString()
			pretty.Println(i, s)
			if !uniqueGenomes[s] {
//				pretty.Println("\tchanging")
				p[j] = self.P[i]
				j++
				uniqueGenomes[s] = true
			}
		} else {
			break
		}
	}
	self.P = p[:j]
}
