package HeapGort
//TODO:figure out why this is super slow
func Ints(slice []int){
  length := len(slice)
  for i := length-1 ; i > 0; i-- {
    for j := (i - 1) / 2; j >= 0; j-- {
      var largest int
      k := j
      for{
        l := 2*k + 1
        r := 2*k + 2
        if l <= i && slice[l] > slice[k] {
          largest = l
        } else {
          largest = k
        }
        if r <= i && slice[r] > slice[largest] {
          largest = r
        }
        if largest != k {
          slice[k], slice[largest] = slice[largest], slice[k]
          k = largest;
        }else{
          break;
        }
      }
    }
    slice[0], slice[i] = slice[i], slice[0]
  }
}
