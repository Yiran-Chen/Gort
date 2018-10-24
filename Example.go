package main
import (
  "./HeapGort"
  "./QuickGort"
  "sort"
  "fmt"
  "time"
  "math/rand"
)
func main() {
  println("Generating pseudo-randeom integers")
  sampleLength := 100000
  rand.Seed(time.Now().Unix())
  a := rand.Perm(sampleLength)
  var algorithms = []string{"QuickGort","Go sort"}
  var t time.Time;
  for i := range algorithms {
    fmt.Println("Running " + algorithms[i]);
    c := make([]int,sampleLength)
    copy(c,a)
    t = time.Now()
    switch algorithms[i] {
      case "HeapGort":
        HeapGort.Ints(c)
      case "QuickGort":
        QuickGort.Ints(c)
      case "Go sort":
        sort.Ints(c)
    }
    if sort.IntsAreSorted(c) {
      fmt.Printf(algorithms[i]+" has been finished in %v\n",time.Since(t))
    }else{
       println("Something wrong in "+algorithms[i])
    }
  }
}
