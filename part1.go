package main
import "fmt"

func konversi(jam,menit,detik int) int{
  return jam * 3600 + menit * 60 + detik
}

func main(){
  var j,m,d int
  fmt.Scanln(&j,&m,&d)
  fmt.Println(konversi(j,m,d))
}
