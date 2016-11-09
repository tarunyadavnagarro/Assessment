package main
import (
	"fmt"
	"regexp"
	"os"
	"bufio"
)
func word_count(words []string) map[string]int {
	
	word_freq := make(map[string]int)
	
	for _, word := range words {
		
		_, ok := word_freq[word]
		
		if ok == true {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}
	return word_freq
}
type word_struct struct {
	freq int
	word string
}

func (p word_struct) String() string {
    return fmt.Sprintf("%3d   %s", p.freq, p.word)
}

func main() {
    
	reader := bufio.NewReader(os.Stdin)

  	fmt.Print("Enter text: ")
  
  	text, _ := reader.ReadString('\n')
 
	
	re := regexp.MustCompile("\\w+")
	words := re.FindAllString(text, -1)
	
    word_map := word_count(words)
    
	
	pws := new(word_struct)
	struct_slice := make([]word_struct, len(word_map))
	ix := 0
	for key, val := range word_map {
		pws.freq = val
		pws.word = key
		
		struct_slice[ix] = *pws
		ix++
	}
	
    for ix := 0; ix < len(struct_slice); ix++ {
        fmt.Printf("%v\n", struct_slice[ix])
    }
   
}

 