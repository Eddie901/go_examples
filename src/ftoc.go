// Ftoc print two Fahrenheit-to-Celsius conversions
package main

import (
	"fmt"
	"tempconv"
)

func main(){
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("[In Kelvin] Brrr! %v\n", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Printf("Freezing point [In Kelvin] Brrr! %v\n", tempconv.CToK(tempconv.FreezingC))
	fmt.Printf("Freezing point [In Celsius] Brrr! %v\n", tempconv.FreezingC)
	fmt.Printf("Freezing point [In Fahrenheit] Brrr! %v\n", tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("Boiling point [In Celsius] Toastie! %v\n", tempconv.BoilingC)
	fmt.Printf("Boiling point [In Fahrenheit] Toastie! %v\n", tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("Boiling point [In Kelvin] Toastie! %v\n", tempconv.CToK(tempconv.BoilingC))
}
