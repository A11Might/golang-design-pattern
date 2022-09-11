package chainofresponsibility

import "fmt"

// Support 处理者(handler)
type Support struct {
	SupportImpl

	name string
	next *Support
}

func NewSupport(impl SupportImpl, name string) *Support {
	return &Support{
		SupportImpl: impl,
		name:        name,
	}
}

func (s *Support) SetNext(next *Support) *Support {
	s.next = next
	return next
}

func (s *Support) Support(trouble *Trouble) {
	if s.resolve(trouble) {
		s.done(trouble)
	} else if s.next != nil {
		s.next.Support(trouble)
	} else {
		s.fail(trouble)
	}
}

func (s *Support) ToString() string {
	return fmt.Sprintf("[%s]", s.name)
}

func (s *Support) done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble.ToString(), s.ToString())
}

func (s *Support) fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble.ToString())
}

type SupportImpl interface {
	resolve(*Trouble) bool
}

// NoSupport 具体的处理者(concrete handler)
type NoSupport struct {
	Support
}

func NewNoSupport(name string) *Support {
	return NewSupport(&NoSupport{}, name)
}

func (n *NoSupport) resolve(trouble *Trouble) bool {
	return false
}

// LimitSupport 具体的处理者(concrete handler)
type LimitSupport struct {
	Support

	limit int
}

func NewLimitSupport(name string, limit int) *Support {
	return NewSupport(&LimitSupport{limit: limit}, name)
}

func (l *LimitSupport) resolve(trouble *Trouble) bool {
	if trouble.GetNumber() < l.limit {
		return true
	} else {
		return false
	}
}

// OddSupport 具体的处理者(concrete handler)
type OddSupport struct {
	Support
}

func NewOddSupport(name string) *Support {
	return NewSupport(&OddSupport{}, name)
}

func (o *OddSupport) resolve(trouble *Trouble) bool {
	if trouble.GetNumber()%2 == 1 {
		return true
	} else {
		return false
	}
}

// SpecialSupport 具体的处理者(concrete handler)
type SpecialSupport struct {
	Support

	number int
}

func NewSpecialSupport(name string, number int) *Support {
	return NewSupport(&SpecialSupport{number: number}, name)
}

func (n *SpecialSupport) resolve(trouble *Trouble) bool {
	if trouble.GetNumber() == n.number {
		return true
	} else {
		return false
	}
}
