package main

import "errors"

type (
	Gene struct {
		Value int
	}

	Chromosome struct {
		Genes []*Gene
	}
)

func NewGene(value int) *Gene {
	return &Gene{value}
}

func NewChromosome(genes []*Gene) *Chromosome {
	return &Chromosome{genes}
}

func PMX(parent1, parent2 *Chromosome) [2]*Chromosome {
	if len(parent1.Genes) != len(parent2.Genes) {
		panic(errors.New("The chromosomes must contain the same number of genes"))
	}

	child1 := new(Chromosome)
	child2 := new(Chromosome)
	children := [2]*Chromosome{child1, child2}

	return children
}

func main() {
}
