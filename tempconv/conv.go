package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kenvin.
func CToK(c Celsius) Kenvin { return Kenvin((c + 273.15)) }

// KToF converts a Kenvin temperature to Celsius.
func KToF(k Celsius) Celsius { return Celsius((k - 273.15)) }
