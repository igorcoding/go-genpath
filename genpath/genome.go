package genpath
import "strconv"

type FitnessT float64
type GenesT []int

type Genome struct {
	conf *GenPathConf
	Fitness FitnessT
	Genes GenesT
}

func NewGenome(conf *GenPathConf) *Genome {
	self := newGenome(conf)

	hasGenes := make(map[int]bool)
	for i := range(self.Genes) {
		if (i == 0) {
			self.Genes[i] = self.conf.StartNode
		} else {
			self.Genes[i] = -1
			for self.Genes[i] == -1 || hasGenes[self.Genes[i]] {
				self.Genes[i] = randInt(0, len(conf.Graph))
			}
		}
		hasGenes[self.Genes[i]] = true
	}
	return self
}

func newGenome(conf *GenPathConf) *Genome {
	self := &Genome{
		conf: conf,
		Fitness: 0,
	}
	self.Genes = make(GenesT, len(conf.Graph))
	return self
}

func (self *Genome) Mutate() *Genome {
	toChange := self.conf.StartNode
	for toChange == self.conf.StartNode {
		toChange = randInt(0, len(self.Genes))
	}
	self.Genes[toChange] = randInt(0, len(self.conf.Graph))
	return self
}

func (self *Genome) Crossover(another *Genome) *Genome {
	g := newGenome(self.conf)
	genomes := []*Genome { self, another }
	activeGenome := randInt(0, 2);
	splitPoints := make(map[int]bool)
	for i := 0; i < self.conf.CrossoverSegmentSplitsCount - 1; i++ {
		s := int(len(self.Genes) / self.conf.CrossoverSegmentSplitsCount) * (i + 1)
		splitPoints[s] = true
	}
	for i := 0; i < len(g.Genes); i++ {
		if splitPoints[i] {
			activeGenome = (activeGenome + 1) % len(genomes)
		}
		g.Genes[i] = genomes[activeGenome].Genes[i]
	}

	return g
}

func (self *Genome) EvalFitness() FitnessT {
	self.Fitness = self.conf.fitnessFunc(self)
	return self.Fitness
}

func (self *Genome) ToString() string {
	var s string
	for i := range(self.Genes) {
		s += strconv.Itoa(self.Genes[i])
	}
	return s
}





type Genomes []*Genome

func (g Genomes) Len() int           { return len(g) }
func (g Genomes) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g Genomes) Less(i, j int) bool {
	return g[i].Fitness > g[j].Fitness
}

func NewGenomes(conf *GenPathConf, count int) Genomes {
	self := make(Genomes, count)
	for i := range(self) {
		self[i] = NewGenome(conf)
	}
	return self
}