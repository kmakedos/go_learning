package temperature

type Celsius float64
type Fahrenheit float64

func CtoF(t Celsius) (f Fahrenheit){
	return Fahrenheit(t * 9/5 + 32)
}

func FtoC(t Fahrenheit) (c Celsius){
	return Celsius((t - 32) * 5 / 9)
}
