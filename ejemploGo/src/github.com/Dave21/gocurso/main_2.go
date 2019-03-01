package main
import "fmt"

const helloworld string = "Hola %s %s, Bienvenido al mundo de Go.\n"
const testConst = "Test"

func main() {
 
  lastname := "Modificar apellido"
  name := getName()
  number := sum(50,50)
  a,b,c := getVariables()
  f32, f64 := getFloat()
 
  stringUTF8:= getUnicode()



  fmt.Printf(helloworld, name, lastname)
  fmt.Println(number,a,b,c)	

  
  fmt.Println(f32,f64)
  fmt.Println("cadena con utf8: " , stringUTF8)

  // con el [0] se obtiene el valor ascii de la letra
  // pero si lo convertimos a string nos dara la letra como tal
  fmt.Println(string("hola"[0]))

  // cantidad de letras que tiene hola
  fmt.Print("Numero de letras que tiene hola ", len("hola"))
} 

func getName() string{
	 var name string
	 name = "sin nombre"
	 fmt.Print("Ingresa tu nombre: ")
   fmt.Scanf("%s", &name)
     return name
}

func getVariables()(int,int32,int64) {
return 1,2,3
}

func getFloat() (float32, float64) {
  return float32(0.1), float64(float32(0.1))
}

func sum(a int, b int) int{	
  var c = a+b
  return c	
}

func getUnicode() string {
  return "こんにちは世界"
}


