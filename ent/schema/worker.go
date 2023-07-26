package schema

import "entgo.io/ent"

// Worker holds the schema definition for the Worker entity.
type Worker struct {
	ent.Schema
}

// Fields of the Worker.
func (Worker) Fields() []ent.Field {
	return nil
}

// Edges of the Worker.
func (Worker) Edges() []ent.Edge {
	return nil
}
