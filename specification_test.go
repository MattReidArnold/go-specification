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

type Trainable interface {
	GetLearnableSkills() []string
}

func (d *Dog) GetLearnableSkills() []string {
	if d == nil {
		return []string{}
	}
	return d.LearnableSkills
}

type Biter interface {
	GetToothSize() float64
}

func (d *Dog) GetToothSize() float64 {
	if d == nil {
		return 0.0
	}
	return d.ToothSize
}

type Purchasable interface {
	GetCost() float64
}

func (d *Dog) GetCost() float64 {
	if d == nil {
		return 0.0
	}
	return d.Cost
}

type PottyTrainable struct {
	specification.Specification
}

func NewPottyTrainable() specification.Specification {
	s := &PottyTrainable{&specification.BaseSpecification{}}
	s.Relate(s)
	return s
}

func (s *PottyTrainable) IsSatisfiedBy(i interface{}) bool {
	t, ok := i.(Trainable)
	if !ok {
		return false
	}
	for _, value := range t.GetLearnableSkills() {
		if value == "house broken" {
			return true
		}
	}
	return false
}

func TestPottyTrainable(t *testing.T) {
	t.Run("PottyTrainable dog", func(t *testing.T) {
		pottyTrainable := NewPottyTrainable()

		d := &Dog{
			LearnableSkills: []string{"sit", "stay", "house broken"},
		}

		want := true
		// true!
		got := pottyTrainable.IsSatisfiedBy(d)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})

	t.Run("Not PottyTrainable dog", func(t *testing.T) {
		notPottyTrainable := NewPottyTrainable().Not()

		d := &Dog{
			LearnableSkills: []string{"lay on couch"},
		}

		want := true
		got := notPottyTrainable.IsSatisfiedBy(d)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})
}

type Dangerous struct {
	specification.Specification
}

func NewDangerous() specification.Specification {
	s := &Dangerous{&specification.BaseSpecification{}}
	s.Relate(s)
	return s
}

func (s *Dangerous) IsSatisfiedBy(i interface{}) bool {
	b, ok := i.(Biter)
	if !ok {
		return false
	}
	return b.GetToothSize() > 3
}

func TestDangerous(t *testing.T) {
	t.Run("Dangerous dog", func(t *testing.T) {
		spec := NewDangerous()

		b := &Dog{
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

		b := &Dog{
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
	s := &Affordable{&specification.BaseSpecification{}}
	s.Relate(s)
	return s
}

func (s *Affordable) IsSatisfiedBy(i interface{}) bool {
	p, ok := i.(Purchasable)
	if !ok {
		return false
	}
	return p.GetCost() < 2000
}

func TestAffordable(t *testing.T) {
	t.Run("Affordable dog", func(t *testing.T) {
		spec := NewAffordable()

		b := &Dog{
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

		b := &Dog{
			Cost: 2000.00,
		}

		want := false
		got := spec.IsSatisfiedBy(b)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}
	})
}

func TestAdoptable(t *testing.T) {
	pottyTrainable := NewPottyTrainable()
	dangerous := NewDangerous()
	affordable := NewAffordable()

	adoptable := pottyTrainable.And(affordable).And(dangerous.Not())

	t.Run("adoptable dog", func(t *testing.T) {

		d := &Dog{
			Cost:            1500.00,
			LearnableSkills: []string{"house broken"},
			ToothSize:       2,
		}

		want := true
		got := adoptable.IsSatisfiedBy(d)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}

	})

	t.Run("too expensive dog", func(t *testing.T) {

		d := &Dog{
			Cost:            2100.00,
			LearnableSkills: []string{"house broken"},
			ToothSize:       2,
		}

		want := false
		got := adoptable.IsSatisfiedBy(d)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}

	})

	t.Run("too dangerous dog", func(t *testing.T) {

		d := &Dog{
			Cost:            1500.00,
			LearnableSkills: []string{"house broken"},
			ToothSize:       5,
		}

		want := false
		got := adoptable.IsSatisfiedBy(d)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}

	})

	t.Run("good for nothing dog", func(t *testing.T) {

		d := &Dog{
			Cost:            1500.00,
			LearnableSkills: []string{},
			ToothSize:       2,
		}

		want := false
		got := adoptable.IsSatisfiedBy(d)

		if got != want {
			t.Errorf("want %v, got %v.\n", want, got)
		}

	})
}
