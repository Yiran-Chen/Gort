package QuickGort

func Ints(slice []int) {
	length := len(slice)
	if length <= 1 {
		return
	}
	i,p := 0,0
	j := length - 1
	key := slice[0]
	for i <= j {
    for j > i && slice[j] >= key {
      j--
    }
    slice[j],slice[p] = slice[p],slice[j];
    p = j
    for i < j && slice[i] <= key {
      i++
    }
    slice[i],slice[p] = slice[p],slice[i];
    p = i
    if i == j {
      break;
    }
	}
  Ints(slice[:p])
  Ints(slice[p+1:])
}
