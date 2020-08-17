package specification

type Specification interface {
	IsSatisfiedBy(interface{}) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}

type BaseSpecification struct {
	Specification
}

// Check specification
func (s *BaseSpecification) IsSatisfiedBy(t interface{}) bool {
	return false
}

// Condition AND
func (s *BaseSpecification) And(spec Specification) Specification {
	a := &AndSpecification{
		s.Specification, spec,
	}
	a.Relate(a)
	return a
}

// Condition OR
func (s *BaseSpecification) Or(spec Specification) Specification {
	a := &OrSpecification{
		s.Specification, spec,
	}
	a.Relate(a)
	return a
}

// Condition NOT
func (s *BaseSpecification) Not() Specification {
	a := &NotSpecification{
		s.Specification,
	}
	a.Relate(a)
	return a
}

// Relate to specification
func (s *BaseSpecification) Relate(spec Specification) {
	s.Specification = spec
}

// AndSpecification
type AndSpecification struct {
	Specification
	compare Specification
}

// Check specification
func (s *AndSpecification) IsSatisfiedBy(t interface{}) bool {
	return s.Specification.IsSatisfiedBy(t) && s.compare.IsSatisfiedBy(t)
}

// OrSpecification
type OrSpecification struct {
	Specification
	compare Specification
}

// Check specification
func (s *OrSpecification) IsSatisfiedBy(t interface{}) bool {
	return s.Specification.IsSatisfiedBy(t) || s.compare.IsSatisfiedBy(t)
}

// NotSpecification
type NotSpecification struct {
	Specification
}

// Check specification
func (s *NotSpecification) IsSatisfiedBy(t interface{}) bool {
	return !s.Specification.IsSatisfiedBy(t)
}
