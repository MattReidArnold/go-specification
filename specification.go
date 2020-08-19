package specification

// Satisfiable interface that allows chaining of specifications
type Satisfiable interface {
	IsSatisfiedBy(interface{}) bool
	And(Satisfiable) Satisfiable
	Or(Satisfiable) Satisfiable
	Not() Satisfiable
	Relate(Satisfiable)
}

// Specification is embedded in all specifications to provide chaining behavior
type Specification struct {
	Satisfiable
}

// IsSatisfiedBy for Specification should never return true. In practice the class
// that has an embedded Specification should override this method.
func (s *Specification) IsSatisfiedBy(t interface{}) bool {
	return false
}

// And combines two Satisfiable Specification using AND logic
func (s *Specification) And(spec Satisfiable) Satisfiable {
	a := &AndSpecification{
		s.Satisfiable, spec,
	}
	a.Relate(a)
	return a
}

// Or combines two Satisfiable Specification using OR logic
func (s *Specification) Or(spec Satisfiable) Satisfiable {
	a := &OrSpecification{
		s.Satisfiable, spec,
	}
	a.Relate(a)
	return a
}

// Not inverts the result of Satisfiable Specification
func (s *Specification) Not() Satisfiable {
	a := &NotSpecification{
		s.Satisfiable,
	}
	a.Relate(a)
	return a
}

// Relate creates a circular embedding between abstract Specification and concrete Specifiable
// This enables functions defined on embedded abstract Specification to be able to call
// IsSatisfiedBy on encloseing concrete specification
func (s *Specification) Relate(spec Satisfiable) {
	s.Satisfiable = spec
}

// AndSpecification combines two Satisfiable specifications with AND logic
type AndSpecification struct {
	Satisfiable
	compare Satisfiable
}

// IsSatisfiedBy verifies both Satisfiable specifications
func (s *AndSpecification) IsSatisfiedBy(t interface{}) bool {
	return s.Satisfiable.IsSatisfiedBy(t) && s.compare.IsSatisfiedBy(t)
}

// OrSpecification combines two Satisfiable specifications with OR logic
type OrSpecification struct {
	Satisfiable
	compare Satisfiable
}

// IsSatisfiedBy verifies at least one Satisfiable specifications
func (s *OrSpecification) IsSatisfiedBy(t interface{}) bool {
	return s.Satisfiable.IsSatisfiedBy(t) || s.compare.IsSatisfiedBy(t)
}

// NotSpecification inverts the logic of a Satisfiable specification
type NotSpecification struct {
	Satisfiable
}

// IsSatisfiedBy verifies a Satisfiable specification is not satisfied
func (s *NotSpecification) IsSatisfiedBy(t interface{}) bool {
	return !s.Satisfiable.IsSatisfiedBy(t)
}
