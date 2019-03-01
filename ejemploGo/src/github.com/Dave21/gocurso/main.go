package main
import "fmt"

const helloworld string = "Hola %s %s, Bienvenido al mundo de Go.\n"
const testConst = "Test"

func main() {
  // 1 parte
  // var name string
  // asignar valor a variable
  // name = "sin nombre"
  // := atajo para crear variables
  
  lastname := "Modificar apellido"

  // 2 parte
   name := getName()

   number := sum(50,50)

a,b,c := getVariables()


// 1 parte
  // var (
		// a = 1
		// b = 2
		// c = 3 
	 //  )

// 1 parte
// // print imprime un mensaje en pantalla
//   fmt.Print("Ingresa tu nombre: ")
// // permite apartir de un formato %s recibe un string, y se lo asigna a una variable  
//   fmt.Scanf("%s", &name)
//   // %s al imprimir este sirve para evitar la concatenacion de las variables 
  
  // printf  permite el recibir variables en una cadena de texto, formatear texto
  // fmt.Printf("Hola %s, bienvenido al fascinante mundo de Go.\n", name)
  fmt.Printf(helloworld, name, lastname)
  
  // Printlnal final de la cadena hace un salto de linea
  fmt.Println(number,a,b,c)	
} 

func getName() string{
	var name string
	name = "sin nombre"

	 fmt.Print("Ingresa tu nombre: ")
  
     fmt.Scanf("%s", &name)

     return name
}

func getVariables()(int,int,int) {
return 1,2,3
}

func sum(a int, b int) int{
	
  var c = a+b
  return c	
}