package main
import "fmt"

func main() {
     
	fmt.Println("Calculadora")
	
     var numero1 float64
     var numero2 float64

	 fmt.Print("Ingresa un numero a procesar: ")  
     fmt.Scanf("%f", &numero1)
     fmt.Print("Ingresa el segundo numero a procesar: ")  
     fmt.Scanf("%f", &numero2)

      suma, resta := operate(numero1,numero2)
     
     fmt.Printf("La suma es %f \n",suma)
     fmt.Printf("La diferencia es %f \n",resta)
}

func operate(a , b float64) (suma float64, resta float64 ) {
	 suma = a+b
     resta= a-b
  return suma, resta
}

