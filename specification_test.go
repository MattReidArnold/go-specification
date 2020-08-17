package specification_test

import (
	"testing"

	"github.com/mattreidarnold/specification"
)

type Dog struct {
	Cost            float64
	ToothSize       float64
	LearnableSkills []string
}

type PottyTrainable struct {
	specification.Specification
}

func NewPottyTrainable() specification.Specification {
	return &PottyTrainable{&specification.BaseSpecification{}}
}

func (s *PottyTrainable) IsSatisfiedBy(i interface{}) bool {
	switch i.(type) {
	case Dog:
		for _, value := range i.(Dog).LearnableSkills {
			if value == "house broken" {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func TestPottyTrainable(t *testing.T) {
	t.Run("PottyTrainable dog", func(t *testing.T) {
		pottyTrainable := NewPottyTrainable()

		b := Dog{
			LearnableSkills: []string{"sit", "stay", "house broken"},
		}

		want := true
		// true!
		got := pottyTrainable.IsSatisfiedBy(b)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})

	t.Run("Not PottyTrainable dog", func(t *testing.T) {
		pottyTrainable := NewPottyTrainable()

		b := Dog{
			LearnableSkills: []string{"lay on couch"},
		}

		want := false
		got := pottyTrainable.IsSatisfiedBy(b)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})
}

type Dangerous struct {
	specification.Specification
}

func NewDangerous() specification.Specification {
	return &Dangerous{&specification.BaseSpecification{}}
}

func (s *Dangerous) IsSatisfiedBy(i interface{}) bool {
	switch i.(type) {
	case Dog:
		return i.(Dog).ToothSize > 3
	default:
		return false
	}
}

func TestDangerous(t *testing.T) {
	t.Run("Dangerous dog", func(t *testing.T) {
		spec := NewDangerous()

		b := Dog{
			ToothSize: 4,
		}

		want := true
		// true!
		got := spec.IsSatisfiedBy(b)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})

	t.Run("Not Dangerous dog", func(t *testing.T) {
		spec := NewPottyTrainable()

		b := Dog{
			ToothSize: 3,
		}

		want := false
		got := spec.IsSatisfiedBy(b)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})
}

type Affordable struct {
	specification.Specification
}

func NewAffordable() specification.Specification {
	return &Affordable{&specification.BaseSpecification{}}
}

func (s *Affordable) IsSatisfiedBy(i interface{}) bool {
	switch i.(type) {
	case Dog:
		return i.(Dog).Cost < 2000
	default:
		return false
	}
}

func TestAffordable(t *testing.T) {
	t.Run("Affordable dog", func(t *testing.T) {
		spec := NewAffordable()

		b := Dog{
			Cost: 1999.99,
		}

		want := true
		// true!
		got := spec.IsSatisfiedBy(b)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})

	t.Run("Not Affordable dog", func(t *testing.T) {
		spec := NewPottyTrainable()

		b := Dog{
			Cost: 2000.00,
		}

		want := false
		got := spec.IsSatisfiedBy(b)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})
}
