package main
import (
  "./Gort"
  "sort"
  "fmt"
   "time"
  "math/rand"
)
func main() {
  println("Generating pseudo-randeom integers");
  sampleLength := 100000
  s := rand.Perm(sampleLength)
  c := make([]int,sampleLength)
  copy(c,s)
  println("Running QuickGort");
  t1 := time.Now()
  QuickGort.Ints(s)
  if(sort.IntsAreSorted(s)){
    fmt.Printf("QuickGort has been finished in %v\n",time.Since(t1));
  }else{
    println("Something wrong in QuickGort");
  }
  println("Running go sort");
  t2 := time.Now()
  sort.Ints(c)
  fmt.Printf("Go sort has been finished in %v\n",time.Since(t2));
}
