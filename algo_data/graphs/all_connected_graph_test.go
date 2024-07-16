package graphs

import (
	"testing"

	"github.com/gpbPiazza/notebook/algo_data/stdout"
	"github.com/stretchr/testify/assert"
)

func makeFriendsNetwork() *cVertex {
	alice := NewCVertex("alice")
	bob := NewCVertex("bob")
	fred := NewCVertex("fred")
	helen := NewCVertex("helen")
	candy := NewCVertex("candy")
	derek := NewCVertex("derek")
	elaine := NewCVertex("elaine")
	gina := NewCVertex("gina")
	irena := NewCVertex("irena")

	alice.addAdjancent(bob)
	alice.addAdjancent(candy)
	alice.addAdjancent(derek)
	alice.addAdjancent(elaine)

	bob.addAdjancent(fred)
	fred.addAdjancent(helen)
	helen.addAdjancent(candy)

	derek.addAdjancent(elaine)
	derek.addAdjancent(gina)
	gina.addAdjancent(irena)
	return alice
}

func Test_cVertex_dfs_transverse(t *testing.T) {
	anVertex := makeFriendsNetwork()
	got := stdout.String(func() { dfs_transverse(anVertex, make(map[*cVertex]bool)) })
	assert.Equal(t, "alice\nbob\nfred\nhelen\ncandy\nderek\nelaine\ngina\nirena\n", got)
}

func Test_cVertex_bfs_transverse(t *testing.T) {
	anVertex := makeFriendsNetwork()
	got := stdout.String(func() { bfs_transverse(anVertex) })
	assert.Equal(t, "alice\nbob\ncandy\nderek\nelaine\nfred\nhelen\ngina\nirena\n", got)
}

func Test_shortestPath(t *testing.T) {
	t.Run("graph from exercise", func(t *testing.T) {
		idris := NewCVertex("Idris")
		talia := NewCVertex("Talia")
		ken := NewCVertex("Ken")
		marco := NewCVertex("Marco")
		sasha := NewCVertex("Sasha")
		lina := NewCVertex("Lina")
		kamil := NewCVertex("Kamil")

		idris.addAdjancent(kamil)
		idris.addAdjancent(talia)
		kamil.addAdjancent(lina)
		lina.addAdjancent(sasha)
		sasha.addAdjancent(marco)
		marco.addAdjancent(ken)
		ken.addAdjancent(talia)

		expectedPath := []string{"Idris", "Kamil", "Lina"}

		shortestPath := shortestPath(idris, lina)

		assert.Equal(t, expectedPath, shortestPath)
	})

	t.Run("test 1", func(t *testing.T) {
		alice := NewCVertex("alice")
		bob := NewCVertex("bob")
		fred := NewCVertex("fred")
		helen := NewCVertex("helen")
		candy := NewCVertex("candy")
		derek := NewCVertex("derek")
		elaine := NewCVertex("elaine")
		gina := NewCVertex("gina")
		irena := NewCVertex("irena")

		alice.addAdjancent(bob)
		alice.addAdjancent(candy)
		alice.addAdjancent(derek)
		alice.addAdjancent(elaine)

		bob.addAdjancent(fred)
		fred.addAdjancent(helen)
		helen.addAdjancent(candy)

		derek.addAdjancent(elaine)
		derek.addAdjancent(gina)
		gina.addAdjancent(irena)

		expectedPath := []string{"alice", "derek", "gina", "irena"}

		result := shortestPath(alice, irena)

		assert.Equal(t, expectedPath, result)
	})
}
