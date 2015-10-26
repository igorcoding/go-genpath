package genpath
import (
	"sort"
	"math/rand"
)

type Population struct {
	conf *GenPathConf
	p Genomes
}

func NewPopulation(conf *GenPathConf) *Population {
	self := &Population{
		conf: conf,
		p: NewGenomes(conf, conf.PopulationSize),
	}

	for i := range(self.p) {
		self.p[i].EvalFitness()
	}
	self.sort()
	return self
}

func (self *Population) evaluate() {
	for k := 0; k < self.conf.CrossoversCount; k++ {
		if rand.Float64() < self.conf.CrossoverProb {
			g1 := self.p[rand.Int31n(int32(self.conf.SelectionCount))]
			g2 := self.p[rand.Int31n(int32(self.conf.SelectionCount))]

			child := g1.Crossover(g2)

			if rand.Float64() < self.conf.MutationProb {
				child.Mutate()
			}
			child.EvalFitness()

			self.p = append(self.p, child)
		}
	}
//	fmt.Println(len(self.p))

	self.sort()
	self.p = self.p[:self.conf.PopulationSize]
//	fmt.Println(len(self.p))
}

func (self *Population) sort() {
	sort.Sort(self.p)
}