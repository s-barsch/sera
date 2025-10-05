package store

import (
	"path/filepath"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/entry/types/tree"
)

const (
	graphDir = "graph"
	cacheDir = "cache"
	aboutDir = "about"
	extraDir = "extra"
)

type Store struct {
	trees
	flats
}

type trees struct {
	graph *tree.Tree
	cache *tree.Tree
	about *tree.Tree
	extra *tree.Tree
}

type flats struct {
	graph entry.Entries
	cache entry.Entries
	about entry.Entries
	extra entry.Entries
}

func (s *Store) TreeByString(section string) (*tree.Tree, bool) {
	switch section {
	case "graph":
		return s.trees.graph, true
	case "cache":
		return s.trees.cache, true
	case "about":
		return s.trees.about, true
	case "extra":
		return s.trees.extra, true
	}

	return nil, false
}

func (s *Store) Graph() *tree.Tree {
	return s.trees.graph
}

func (s *Store) Cache() *tree.Tree {
	return s.trees.cache
}

func (s *Store) About() *tree.Tree {
	return s.trees.about
}

func (s *Store) Extra() *tree.Tree {
	return s.trees.extra
}

func (s *Store) GraphFlat() entry.Entries {
	return s.flats.graph
}

func (s *Store) CacheFlat() entry.Entries {
	return s.flats.cache
}

func (s *Store) AboutFlat() entry.Entries {
	return s.flats.about
}

func (s *Store) ExtraFlat() entry.Entries {
	return s.flats.extra
}

func ReadPublic(root string) (*Store, error) {
	store, err := ReadPrivate(root)
	if err != nil {
		return nil, err
	}

	trees := trees{
		graph: store.Graph().Public(),
		cache: store.Cache().Public(),
		about: store.About().Public(),
		extra: store.Extra().Public(),
	}

	return &Store{
		trees: trees,
		flats: flats{
			graph: trees.graph.Flatten(),
			cache: trees.cache.Flatten(),
			about: trees.about.Flatten(),
			extra: trees.extra.Flatten(),
		},
	}, nil
}

func ReadPrivate(root string) (*Store, error) {
	graph, err := tree.ReadTree(filepath.Join(root, graphDir), nil)
	if err != nil {
		return nil, err
	}

	cache, err := tree.ReadTree(filepath.Join(root, cacheDir), nil)
	if err != nil {
		return nil, err
	}

	about, err := tree.ReadTree(filepath.Join(root, aboutDir), nil)
	if err != nil {
		return nil, err
	}

	extra, err := tree.ReadTree(filepath.Join(root, extraDir), nil)
	if err != nil {
		return nil, err
	}

	return &Store{
		trees: trees{
			graph: graph,
			cache: cache,
			about: about,
			extra: extra,
		},
		flats: flats{
			graph: graph.Flatten(),
			cache: cache.Flatten(),
			about: about.Flatten(),
			extra: extra.Flatten(),
		},
	}, nil
}
