// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CtoK(c Celsius) Kelvin { return Kelvin(c - 273.15) }
func KtoC(c Kelvin) Celsius { return Celsius(c + 273.15) }
