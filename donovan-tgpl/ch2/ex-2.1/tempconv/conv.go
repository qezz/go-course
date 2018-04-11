package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) } // do we handle negative Kelvin values?

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) } // meaningless hack

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(Celsius(k) + AbsoluteZeroC) }

// KToC converts a Kelvin temperature to Fahrenheit.
func KtoF(k Kelvin) Fahrenheit { return CToF(KToC(k)) } // another one
