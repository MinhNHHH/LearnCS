package exercises

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }
type fahrenheitFlag struct{ Fahrenheit }

// Implement the String method for celsiusFlag
func (cf *celsiusFlag) String() string {
	return fmt.Sprintf("%v°C", cf.Celsius)
}

func (cf *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		cf.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		cf.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (ff *fahrenheitFlag) String() string {
	return fmt.Sprintf("%vF", ff.Fahrenheit)
}

func (ff *fahrenheitFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		ff.Fahrenheit = CToF(Celsius(value))
		return nil
	case "F", "°F":
		ff.Fahrenheit = Fahrenheit(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func KevinFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	f := fahrenheitFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Fahrenheit
}
