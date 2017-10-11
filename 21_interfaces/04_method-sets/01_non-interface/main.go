/*
https://github.com/golang/go/wiki/MethodSets

Spec1 - Method Sets:
A type may have a method set associated with it. The method set of an interface type is its interface.
The method set of any other named type T consists of all methods with receiver type T.
The method set of the corresponding pointer type *T is the set of all methods with receiver *T or T
(that is, it also contains the method set of T).
Any other type has an empty method set.
In a method set, each method must have a unique name.

Spec2 - Calls:
A method call x.m() is valid if the method set of (the type of) x contains m and
the argument list can be assigned to the parameter list of m.
If x is addressable and &x's method set contains m, x.m() is shorthand for (&x).m().
*/
package main

import (
	"fmt"
)

// List type is type of an int slice
type List []int

// Len returns the length of the List.
func (l List) Len() int { return len(l) }

// Append val to the List.
func (l *List) Append(val int) { *l = append(*l, val) }

func main() {

	// A bare value
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len()) // [1] (len: 1)

	// A pointer value
	plst := new(List)
	plst.Append(2)
	// we can call Len on a pointer value plst.
	fmt.Printf("%v (len: %d)\n", plst, plst.Len()) // &[2] (len: 1)
}

/*
Note that both pointer and value methods can both be called on both pointer and non-pointer values.
To understand why, let's examine the method sets of both types, directly from the spec:

	List
	- Len() int

	*List
	- Len() int
	- Append(int)

Notice that the method set for List does not actually contain Append(int)
even though you can see from the above program that you can call the method without a problem.
This is a result of the second spec section above.
It implicitly translates the first line below into the second:

	lst.Append(1)
	(&lst).Append(1)

Now that the value before the dot is a *List, its method set includes Append, and the call is legal.

To make it easier to remember these rules,
it may be helpful to simply consider the pointer- and value-receiver methods separately from the method set.

*It is legal to call a pointer-valued method on anything that is already a pointer or
whose address can be taken (as is the case in the above example).

*It is legal to call a value method on anything which is a value or
whose value can be dereferenced (as is the case with any pointer; this case is specified explicitly in the spec).

Values          Receivers
-----------------------------------------------
T               (t T) and (t *T)
*T              (t T) and (t *T)


*/
