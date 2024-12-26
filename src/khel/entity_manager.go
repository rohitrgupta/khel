package khel

type EntityManager struct {
	Entities []*Entity
	addQueue []*Entity
	tags     map[string][]*Entity
	size     int
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities: make([]*Entity, 0),
		addQueue: make([]*Entity, 0),
		tags:     make(map[string][]*Entity),
		size:     0,
	}
}

func (em *EntityManager) AddEntity(tag string) *Entity {
	em.size++
	e := NewEntity(em.size, tag)
	em.addQueue = append(em.addQueue, e)
	return e
}

func (em *EntityManager) Update() {
	// add new entities
	for _, e := range em.addQueue {
		em.Entities = append(em.Entities, e)
		em.tags[e.Tag()] = append(em.tags[e.Tag()], e)
	}
	em.addQueue = em.addQueue[:0]
	// remove dead entities
	for i, e := range em.Entities {
		if !e.IsAlive() {
			em.Entities = append(em.Entities[:i], em.Entities[i+1:]...)
			i--
		}
	}
	// remove dead entities from tags
	for tag, entities := range em.tags {
		for i, e := range entities {
			if !e.IsAlive() {
				em.tags[tag] = append(em.tags[tag][:i], em.tags[tag][i+1:]...)
				i--
			}
		}
	}
}

func (em *EntityManager) GetEntityByTag(tag string) []*Entity {
	return em.tags[tag]
}
