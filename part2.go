package main
import "fmt"

func check(a,b,c int) string{
  var cek string 
  if (a+b>c) && (a+c>b) && (b+c>a) {
    cek = "segitiga"
  } else {
    cek = "bukan segitiga"
  }
  return cek
}

func main(){
  var a,b,c int
  fmt.Scanln(&a,&b,&c)
  fmt.Println(check(a,b,c))
}
