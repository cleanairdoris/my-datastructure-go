package datastructure

import (
	"errors"
)

//Stack defined as slice
type Stack []interface{}

//Len return the len of the stack
func (stack Stack) Len() int {
	return len(stack)
}

//Cap return the cap of the stack
func (stack Stack) Cap() int {
	return cap(stack)
}

//Push can add a value into the stack
func (stack *Stack) Push(val interface{}) error {
	if stack == nil {
		return errors.New("the stack is nil")
	}
	*stack = append(*stack, val)
	return nil
}

//Top can get the top value of the stack
func (stack *Stack) Top() (interface{}, error) {
	if stack == nil {
		return nil, errors.New("the stack is nil")
	}
	l := len(*stack)
	if l == 0 {
		return nil, errors.New("the stack len is 0")
	}

	return (*stack)[l-1], nil
}

//Pop can get the top value of the stack and remove it
func (stack *Stack) Pop() (interface{}, error) {
	if stack == nil {
		return nil, errors.New("the stack is nil")
	}
	l := len(*stack)
	if l == 0 {
		return nil, errors.New("the stack len is 0")
	}
	ret := (*stack)[l-1]
	*stack = (*stack)[:l-1]

	return ret, nil
}

//IsEmpty judge if the stack is empty
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
