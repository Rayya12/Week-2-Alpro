package main
import "fmt"

func cek_diskon(total int ,ismember bool) int{
  if (total > 100000 && ismember) {
    return total  * 90 /100
  } else if (total > 100000 && !ismember) {
    return total * 95 /100
  } else {
    return total
  }
}

func main(){
  var total int
  var ismember bool
  fmt.Scanln(&total,&ismember)
  fmt.Println(cek_diskon(total,ismember))
}
