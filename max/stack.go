package max

import Error "github.com/deepset/most-active-cookie/errors"

type Record struct {
	id    string
	count int
}

type stack struct {
	top  int
	list []Record
}

func (s *stack) push(r Record) {
	s.top += 1
	s.list[s.top] = r
}

/* commenting unused pop function
func (s *stack) pop(r Record) error {
	if s.isEmpty() {
		return Error.ErrEmptyStack
	}
	s.list = s.list[:len(s.list)-1]
	s.top -= 1
	return nil
}
*/

func (s *stack) peek() (Record, error) {
	if s.isEmpty() {
		return Record{}, Error.ErrEmptyStack
	}
	record := s.list[s.top]
	return record, nil
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (stack *stack) setEmpty() {
	stack.top = -1
}
