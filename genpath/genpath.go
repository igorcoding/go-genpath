package genpath
import (
	"math"
)

type GenPath struct {
	Conf *GenPathConf

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

	self.Conf = conf
	self.generation = 0
	self.population = NewPopulation(conf)

	self.Inf = FitnessT(math.Inf(0))
	return self, nil
}

func (self *GenPath) Step() (Genomes, error) {
	self.population.evaluate()
	self.generation++
	return self.population.P, nil
}




func (self *GenPath) calcFitness(g *Genome) FitnessT {
	if (len(g.Genes) > 0) {
		if (len(g.Genes) == 1 && self.Conf.StartNode != self.Conf.EndNode) {
			return FitnessT(0)
		}
		if g.Genes[0] != self.Conf.StartNode {
			return FitnessT(0)
		}

		dist := 0.0
		reachedEnd := false
		for i := 0; i < len(g.Genes) - 1; i++ {
			dist += self.Conf.dist(g.Genes[i], g.Genes[i + 1])
			if g.Genes[i + 1] == self.Conf.EndNode {
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