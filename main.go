package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type (
	Gene struct {
		Value int
	}

	Chromosome struct {
		Genes []Gene
	}

	PMX struct {
		Parent1, Parent2 *Chromosome
		StartPoint       int
		EndPoint         int
		Size             int
	}
)

func NewGene(value int) Gene {
	return Gene{value}
}

func NewChromosome(genes []Gene) *Chromosome {
	return &Chromosome{genes}
}

func NewPMX(parent1, parent2 *Chromosome) *PMX {
	chromosomeSize := len(parent1.Genes)
	if len(parent1.Genes) != len(parent2.Genes) {
		panic(errors.New("The chromosomes must contain the same number of genes"))
	}

	startPoint, endPoint := GetCrossoverPoints(chromosomeSize)
	return &PMX{parent1, parent2, startPoint, endPoint, chromosomeSize}
}

func (pmx *PMX) Run() [2]*Chromosome {
	child1Genes := make([]Gene, pmx.Size)
	child2Genes := make([]Gene, pmx.Size)
	for i := pmx.StartPoint; i <= pmx.EndPoint; i++ {
		child1Genes[i] = pmx.Parent2.Genes[i]
		child2Genes[i] = pmx.Parent1.Genes[i]
	}

	fillRemainingGenes(child1Genes, pmx.Parent1.Genes, pmx.StartPoint, pmx.EndPoint)
	fillRemainingGenes(child2Genes, pmx.Parent2.Genes, pmx.StartPoint, pmx.EndPoint)

	children := [2]*Chromosome{NewChromosome(child1Genes), NewChromosome(child2Genes)}
	return children
}

func fillRemainingGenes(childGenes, parentGenes []Gene, startPoint, endPoint int) {
	chromosomeSize := len(parentGenes)
	for i := range chromosomeSize {
		if i >= startPoint && i <= endPoint {
			continue
		}

		gene := parentGenes[i]
		for contains(childGenes, gene) {
			gene = parentGenes[indexOfGene(childGenes, gene)]
		}
		childGenes[i] = NewGene(gene.Value)
	}
}

func contains(genes []Gene, gene Gene) bool {
	for _, g := range genes {
		if g.Value == gene.Value {
			return true
		}
	}
	return false
}

func indexOfGene(genes []Gene, gene Gene) int {
	for i, g := range genes {
		if g.Value == gene.Value {
			return i
		}
	}
	return -1
}

func GetCrossoverPoints(size int) (startPoint, endPoint int) {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	startPoint = rnd.Intn(size)
	endPoint = rnd.Intn(size - 1)
	if endPoint >= startPoint {
		endPoint++
	} else {
		startPoint, endPoint = endPoint, startPoint
	}
	return
}

func main() {
	parent1 := NewChromosome([]Gene{
		NewGene(1),
		NewGene(2),
		NewGene(3),
		NewGene(4),
		NewGene(5),
		NewGene(6),
		NewGene(7),
		NewGene(8),
		NewGene(9),
	})
	parent2 := NewChromosome([]Gene{
		NewGene(4),
		NewGene(5),
		NewGene(6),
		NewGene(9),
		NewGene(8),
		NewGene(1),
		NewGene(3),
		NewGene(2),
		NewGene(7),
	})
	pmx := NewPMX(parent1, parent2)
	children := pmx.Run()
	fmt.Println("Crossover Points =>", "Start:", pmx.StartPoint, "End:", pmx.EndPoint)
	fmt.Println("Parents:")
	printGenes(parent1.Genes)
	printGenes(parent2.Genes)
	fmt.Println("\nChildren:")
	for _, child := range children {
		printGenes(child.Genes)
	}
}

func printGenes(g []Gene) {
	values := make([]int, len(g))
	for i, gene := range g {
		values[i] = gene.Value
	}
	fmt.Println(fmt.Sprint(values))
}
