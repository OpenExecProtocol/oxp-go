// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsToolCallResponseResultObjectValueUnion() {}

type UnionBool bool

func (UnionBool) ImplementsToolCallResponseResultObjectValueUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsToolCallResponseResultObjectValueUnion() {}
