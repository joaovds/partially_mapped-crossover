package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPMX(t *testing.T) {
	t.Run("should return different chromosome size error", func(t *testing.T) {
		parent1Genes := []Gene{
			NewGene(1),
			NewGene(2),
			NewGene(3),
		}
		parent2Genes := []Gene{
			NewGene(2),
			NewGene(3),
		}
		parent1 := NewChromosome(parent1Genes)
		parent2 := NewChromosome(parent2Genes)

		assert.PanicsWithError(t, "The chromosomes must contain the same number of genes", func() {
			NewPMX(parent1, parent2)
		})
	})

	t.Run("should return the crossed children with the correct values", func(t *testing.T) {
		t.Run("1", func(t *testing.T) {
			parent1Genes := []Gene{
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
			parent2Genes := []Gene{
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
			pmx := NewPMX(parent1, parent2)
			// mock
			pmx.StartPoint, pmx.EndPoint = 2, 5

			children := pmx.Run()
			child1 := children[0]
			child2 := children[1]

			t.Run("child1 values", func(t *testing.T) {
				assert.Equal(t, 3, child1.Genes[0].Value)
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

		t.Run("2", func(t *testing.T) {
			parent1Genes := []Gene{
				NewGene(9),
				NewGene(2),
				NewGene(7),
				NewGene(5),
				NewGene(4),
				NewGene(3),
				NewGene(6),
				NewGene(1),
				NewGene(8),
			}
			parent2Genes := []Gene{
				NewGene(2),
				NewGene(8),
				NewGene(3),
				NewGene(6),
				NewGene(9),
				NewGene(5),
				NewGene(7),
				NewGene(4),
				NewGene(1),
			}
			parent1 := NewChromosome(parent1Genes)
			parent2 := NewChromosome(parent2Genes)
			pmx := NewPMX(parent1, parent2)
			// mock
			pmx.StartPoint, pmx.EndPoint = 3, 5

			children := pmx.Run()
			child1 := children[0]
			child2 := children[1]

			t.Run("child1 values", func(t *testing.T) {
				assert.Equal(t, 4, child1.Genes[0].Value)
				assert.Equal(t, 2, child1.Genes[1].Value)
				assert.Equal(t, 7, child1.Genes[2].Value)
				assert.Equal(t, 6, child1.Genes[3].Value)
				assert.Equal(t, 9, child1.Genes[4].Value)
				assert.Equal(t, 5, child1.Genes[5].Value)
				assert.Equal(t, 3, child1.Genes[6].Value)
				assert.Equal(t, 1, child1.Genes[7].Value)
				assert.Equal(t, 8, child1.Genes[8].Value)
			})

			t.Run("child2 values", func(t *testing.T) {
				assert.Equal(t, 2, child2.Genes[0].Value)
				assert.Equal(t, 8, child2.Genes[1].Value)
				assert.Equal(t, 6, child2.Genes[2].Value)
				assert.Equal(t, 5, child2.Genes[3].Value)
				assert.Equal(t, 4, child2.Genes[4].Value)
				assert.Equal(t, 3, child2.Genes[5].Value)
				assert.Equal(t, 7, child2.Genes[6].Value)
				assert.Equal(t, 9, child2.Genes[7].Value)
				assert.Equal(t, 1, child2.Genes[8].Value)
			})
		})

		t.Run("3", func(t *testing.T) {
			parent1Genes := []Gene{
				NewGene(1),
				NewGene(2),
				NewGene(3),
				NewGene(4),
			}
			parent2Genes := []Gene{
				NewGene(2),
				NewGene(3),
				NewGene(1),
				NewGene(4),
			}
			parent1 := NewChromosome(parent1Genes)
			parent2 := NewChromosome(parent2Genes)
			pmx := NewPMX(parent1, parent2)
			// mock
			pmx.StartPoint, pmx.EndPoint = 0, 1

			children := pmx.Run()
			child1 := children[0]
			child2 := children[1]

			t.Run("child1 values", func(t *testing.T) {
				assert.Equal(t, 2, child1.Genes[0].Value)
				assert.Equal(t, 3, child1.Genes[1].Value)
				assert.Equal(t, 1, child1.Genes[2].Value)
				assert.Equal(t, 4, child1.Genes[3].Value)
			})

			t.Run("child2 values", func(t *testing.T) {
				assert.Equal(t, 1, child2.Genes[0].Value)
				assert.Equal(t, 2, child2.Genes[1].Value)
				assert.Equal(t, 3, child2.Genes[2].Value)
				assert.Equal(t, 4, child2.Genes[3].Value)
			})
		})
	})
}

func TestContains(t *testing.T) {
	genes := []Gene{
		NewGene(1),
		NewGene(4),
		NewGene(16),
		NewGene(28),
		NewGene(10),
	}
	existsGene := NewGene(28)
	noExistsGene := NewGene(282)

	assert.True(t, contains(genes, existsGene))
	assert.False(t, contains(genes, noExistsGene))
}

func TestIndexOfGene(t *testing.T) {
	genes := []Gene{
		NewGene(1),
		NewGene(4),
		NewGene(16),
		NewGene(28),
		NewGene(10),
	}

	assert.Equal(t, 3, indexOfGene(genes, NewGene(28)))
	assert.Equal(t, 1, indexOfGene(genes, NewGene(4)))
}
