package main
import "fmt"

func konversi(jam,menit,detik int) int{
  return jam * 3600 + menit * 60 + detik
}

func main(){
  fmt.Println(konversi(0,3,10))
  fmt.Println(konversi(1,0,5))
}
