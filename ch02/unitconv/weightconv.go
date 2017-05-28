// Package unitconv performs US and Metric Weight conversions.
package unitconv

import "fmt"

type Pound float64
type Ounce float64
type Kilogram float64
type Gram float64
type Milligram float64

func (p Pound) String() string     { return fmt.Sprintf("%g lb", p) }
func (o Ounce) String() string     { return fmt.Sprintf("%g oz", o) }
func (k Kilogram) String() string  { return fmt.Sprintf("%g kg", k) }
func (g Gram) String() string      { return fmt.Sprintf("%g g", g) }
func (m Milligram) String() string { return fmt.Sprintf("%g mg", m) }

// PToO converts pounds to ounces
func PToO(p Pound) Ounce { return Ounce(p * 16) }

// OToP converts ounces to pounds
func OToP(o Ounce) Pound { return Pound(o / 16) }

// PToK converts pounds to kilograms
func PToK(p Pound) Kilogram { return Kilogram(p / 2.54) }

// KToP converts kilograms to pounds
func KToP(k Kilogram) Pound { return Pound(k * 2.54) }
