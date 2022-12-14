// Code generated by ent, DO NOT EDIT.

package water

import (
	"minimal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Liters applies equality check predicate on the "liters" field. It's identical to LitersEQ.
func Liters(v float64) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLiters), v))
	})
}

// Topic applies equality check predicate on the "topic" field. It's identical to TopicEQ.
func Topic(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopic), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// LitersEQ applies the EQ predicate on the "liters" field.
func LitersEQ(v float64) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLiters), v))
	})
}

// LitersNEQ applies the NEQ predicate on the "liters" field.
func LitersNEQ(v float64) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLiters), v))
	})
}

// LitersIn applies the In predicate on the "liters" field.
func LitersIn(vs ...float64) predicate.Water {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLiters), v...))
	})
}

// LitersNotIn applies the NotIn predicate on the "liters" field.
func LitersNotIn(vs ...float64) predicate.Water {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLiters), v...))
	})
}

// LitersGT applies the GT predicate on the "liters" field.
func LitersGT(v float64) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLiters), v))
	})
}

// LitersGTE applies the GTE predicate on the "liters" field.
func LitersGTE(v float64) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLiters), v))
	})
}

// LitersLT applies the LT predicate on the "liters" field.
func LitersLT(v float64) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLiters), v))
	})
}

// LitersLTE applies the LTE predicate on the "liters" field.
func LitersLTE(v float64) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLiters), v))
	})
}

// TopicEQ applies the EQ predicate on the "topic" field.
func TopicEQ(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopic), v))
	})
}

// TopicNEQ applies the NEQ predicate on the "topic" field.
func TopicNEQ(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTopic), v))
	})
}

// TopicIn applies the In predicate on the "topic" field.
func TopicIn(vs ...string) predicate.Water {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTopic), v...))
	})
}

// TopicNotIn applies the NotIn predicate on the "topic" field.
func TopicNotIn(vs ...string) predicate.Water {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTopic), v...))
	})
}

// TopicGT applies the GT predicate on the "topic" field.
func TopicGT(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTopic), v))
	})
}

// TopicGTE applies the GTE predicate on the "topic" field.
func TopicGTE(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTopic), v))
	})
}

// TopicLT applies the LT predicate on the "topic" field.
func TopicLT(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTopic), v))
	})
}

// TopicLTE applies the LTE predicate on the "topic" field.
func TopicLTE(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTopic), v))
	})
}

// TopicContains applies the Contains predicate on the "topic" field.
func TopicContains(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTopic), v))
	})
}

// TopicHasPrefix applies the HasPrefix predicate on the "topic" field.
func TopicHasPrefix(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTopic), v))
	})
}

// TopicHasSuffix applies the HasSuffix predicate on the "topic" field.
func TopicHasSuffix(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTopic), v))
	})
}

// TopicEqualFold applies the EqualFold predicate on the "topic" field.
func TopicEqualFold(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTopic), v))
	})
}

// TopicContainsFold applies the ContainsFold predicate on the "topic" field.
func TopicContainsFold(v string) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTopic), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Water {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Water {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Water) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Water) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Water) predicate.Water {
	return predicate.Water(func(s *sql.Selector) {
		p(s.Not())
	})
}
