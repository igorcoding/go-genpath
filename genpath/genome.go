package genpath

type FitnessT float64
type GenesT []int

type Genome struct {
	conf *GenPathConf
	fitness FitnessT
	genes GenesT
}

func NewGenome(conf *GenPathConf) *Genome {
	self := newGenome(conf)

	hasGenes := make(map[int]bool)
	for i := range(self.genes) {
		if (i == 0) {
			self.genes[i] = self.conf.StartNode
		} else {
			self.genes[i] = -1
			for self.genes[i] == -1 || hasGenes[self.genes[i]] {
				self.genes[i] = randInt(0, len(conf.Graph))
			}
		}
		hasGenes[self.genes[i]] = true
	}
	return self
}

func newGenome(conf *GenPathConf) *Genome {
	self := &Genome{
		conf: conf,
		fitness: 0,
	}
	self.genes = make(GenesT, len(conf.Graph))
	return self
}

func (self *Genome) Mutate() *Genome {
	toChange := self.conf.StartNode
	for toChange == self.conf.StartNode {
		toChange = randInt(0, len(self.genes))
	}
	self.genes[toChange] = randInt(0, len(self.conf.Graph))
	return self
}

func (self *Genome) Crossover(another *Genome) *Genome {
	g := newGenome(self.conf)
	genomes := []*Genome { self, another }
	activeGenome := randInt(0, 2);
	splitPoint := int(len(self.genes) / 2)
	for i := 0; i < len(g.genes); i++ {
		if i == splitPoint {
			activeGenome = (activeGenome + 1) % len(genomes)
		}
		g.genes[i] = genomes[activeGenome].genes[i]
	}

	return g
}

func (self *Genome) EvalFitness() FitnessT {
	self.fitness = self.conf.fitnessFunc(self)
	return self.fitness
}





type Genomes []*Genome

func (g Genomes) Len() int           { return len(g) }
func (g Genomes) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g Genomes) Less(i, j int) bool {
	return g[i].fitness > g[j].fitness
}

func NewGenomes(conf *GenPathConf, count int) Genomes {
	self := make(Genomes, count)
	for i := range(self) {
		self[i] = NewGenome(conf)
	}
	return self
}