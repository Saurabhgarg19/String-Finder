package stringfinder

const (
	MaxArraySize int = 100
)

//KMP Return index list if target string contain want string
func KMP(mainStr string, subStr string) []int {
	next := preKMP(mainStr)
	n := len(subStr)
	m := len(mainStr)
  
	y := []byte(subStr)
	x := []byte(mainStr)
	var ret []int

  i := 0
	j := 0
  
	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		if i >= m {
			ret = append(ret, j-i)
			i = next[i]
		}
	}

	return ret
}


func preKMP(x string) [MaxArraySize]int {
    	var i, j int
    	length := len(x) - 1
    	var kmpNext [MaxArraySize]int
    	i = 0
    	j = -1
    	kmpNext[0] = -1
    
    	for i < length {
    		for j > -1 && x[i] != x[j] {
    			j = kmpNext[j]
    	}
    
    	i++
    	j++
    
    	if x[i] == x[j] {
    		kmpNext[i] = kmpNext[j]
    	} else {
    		kmpNext[i] = j
    	}
	}
	return kmpNext
}
