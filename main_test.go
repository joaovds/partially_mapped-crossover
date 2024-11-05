package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPMX(t *testing.T) {
	t.Run("should return different chromosome size error", func(t *testing.T) {
		parent1Genes := []*Gene{
			NewGene(1),
			NewGene(2),
			NewGene(3),
		}
		parent2Genes := []*Gene{
			NewGene(2),
			NewGene(3),
		}
		parent1 := NewChromosome(parent1Genes)
		parent2 := NewChromosome(parent2Genes)

		assert.PanicsWithError(t, "The chromosomes must contain the same number of genes", func() {
			PMX(parent1, parent2)
		})
	})

	t.Run("should return the crossed children with the correct values", func(t *testing.T) {
		parent1Genes := []*Gene{
			NewGene(1),
			NewGene(2),
			NewGene(3),
			NewGene(4),
			NewGene(5),
			NewGene(6),
			NewGene(7),
			NewGene(8),
			NewGene(9),
		}
		parent2Genes := []*Gene{
			NewGene(5),
			NewGene(4),
			NewGene(6),
			NewGene(9),
			NewGene(2),
			NewGene(1),
			NewGene(7),
			NewGene(8),
			NewGene(3),
		}
		parent1 := NewChromosome(parent1Genes)
		parent2 := NewChromosome(parent2Genes)

		children := PMX(parent1, parent2)
		child1 := children[0]
		child2 := children[1]

		t.Run("child1 values", func(t *testing.T) {
			assert.Equal(t, 1, child1.Genes[0].Value)
			assert.Equal(t, 5, child1.Genes[1].Value)
			assert.Equal(t, 6, child1.Genes[2].Value)
			assert.Equal(t, 9, child1.Genes[3].Value)
			assert.Equal(t, 2, child1.Genes[4].Value)
			assert.Equal(t, 1, child1.Genes[5].Value)
			assert.Equal(t, 7, child1.Genes[6].Value)
			assert.Equal(t, 8, child1.Genes[7].Value)
			assert.Equal(t, 4, child1.Genes[8].Value)
		})

		t.Run("child2 values", func(t *testing.T) {
			assert.Equal(t, 2, child2.Genes[0].Value)
			assert.Equal(t, 9, child2.Genes[1].Value)
			assert.Equal(t, 3, child2.Genes[2].Value)
			assert.Equal(t, 4, child2.Genes[3].Value)
			assert.Equal(t, 5, child2.Genes[4].Value)
			assert.Equal(t, 6, child2.Genes[5].Value)
			assert.Equal(t, 7, child2.Genes[6].Value)
			assert.Equal(t, 8, child2.Genes[7].Value)
			assert.Equal(t, 1, child2.Genes[8].Value)
		})
	})
}
