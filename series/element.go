package series

import "fmt"

// Type is a convenience alias that can be used for a more type safe way of reason and use Series types.
type ElementType string

// Deprecated: Type will be removed; use ElementType instead
type Type = ElementType

// ElementValue represents the value that can be used for marshaling or unmarshaling Elements.
type ElementValue interface{}

// Element is the interface that defines the types of methods to be present for elements of a Series
type Element interface {
	Set(interface{}) // Deprecated: Set is not neccessary as an Element will become immutable

	// Comparation methods

	Eq(Element) bool // Equals
	Ne(Element) bool // Not Equal
	Lt(Element) bool // Less Than
	Le(Element) bool // Less or Equal
	Gt(Element) bool // Greater Than
	Ge(Element) bool // Greater or Equal

	Neq(Element) bool       // Deprecated: use Ne instead
	Less(Element) bool      // Deprecated: use Lt instead
	LessEq(Element) bool    // Deprecated: use Le instead
	Greater(Element) bool   // Deprecated: use Gt instead
	GreaterEq(Element) bool // Deprecated: use Ge instead

	// Accessor/conversion methods
	Copy() Element     // Deprecated: Copy is not neccassary as an Element will become immutable
	Val() ElementValue // FIXME: Returning interface is a recipe for pain

	Int() (int, error)
	Float() float64
	Bool() (bool, error)

	// Information methods
	IsNA() bool
	Type() ElementType

	fmt.Stringer
}
