// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(k Kelvin) Kelvin { return Kelvin(Celsius((k-32)*5/9) + AbsoluteZeroC) }

// CToK converts a Fahrenheit temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

// KtoF converts a Fahrenheit temperature to Kelvin.
func KtoF(k Kelvin) Fahrenheit { return CToF(Celsius(k - AbsoluteZeroK)) }

// KtoC converts a Fahrenheit temperature to Kelvin.
func KtoC(k Kelvin) Celsius { return Celsius(k + AbsoluteZeroK) }

//!-
