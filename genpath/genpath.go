package genpath
import (
	"math"
)

type GenPath struct {
	conf *GenPathConf

	generation int
	population *Population

	Inf FitnessT
}

func NewGenPath(conf *GenPathConf) (*GenPath, error) {
	self := &GenPath{}
	conf.fitnessFunc = self.calcFitness
	err := conf.check()
	if err != nil {
		return nil, err
	}

	self.conf = conf
	self.generation = 0
	self.population = NewPopulation(conf)

	self.Inf = FitnessT(math.Inf(0))
	return self, nil
}

func (self *GenPath) Step() error {
	self.population.evaluate()
	self.generation++
	return nil
}




func (self *GenPath) calcFitness(g *Genome) FitnessT {
	if (len(g.genes) > 0) {
		if (len(g.genes) == 1 && self.conf.StartNode != self.conf.EndNode) {
			return FitnessT(0)
		}
		if g.genes[0] != self.conf.StartNode {
			return FitnessT(0)
		}

		dist := 0.0
		reachedEnd := false
		for i := 0; i < len(g.genes) - 1; i++ {
			dist += self.conf.dist(g.genes[i], g.genes[i + 1])
			if g.genes[i + 1] == self.conf.EndNode {
				reachedEnd = true
				break
			}
		}
		if !reachedEnd {
			dist = float64(self.Inf)
		}
		if dist == 0 {
			return self.Inf
		}
		return FitnessT(1.0 / dist)
	}
	return FitnessT(0)
}