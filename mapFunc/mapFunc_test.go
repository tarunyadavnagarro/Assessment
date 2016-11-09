package main
import (
"testing"
"reflect"
)
func TestmapFunc(t *testing.T) {

	
	input:=[]string{}
	map_val := word_count(input)
	
	if map_val!=nil	{
		t.Fatal("failed for nil ")
	}

	input =[]string{"bat","go","cat","gopher", "go"}

	output:=map[string]int{"bat" :1,"go" :2,"cat" :1,"gopher" :1,}

	val:=word_count(input)
  		if(reflect.DeepEqual(output,val)==false){
          t.Fatal("failed")
      	}

}
