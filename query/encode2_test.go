package query

import (
	"fmt"
	"net/url"
	"testing"
)

type OtherProps struct {
	Height string
	Weight int
}

type Person struct {
	Name       string
	Age        int
	KidsNames  []string
	SomeProps  map[string]interface{}
	OtherProps *OtherProps
}

func Test2(t *testing.T) {
	return
	p := Person{"Drew", 30, []string{"Tom", "Buck", "Fred"}, map[string]interface{}{
		"Hair":    "Brown",
		"Fingers": 10,
	}, &OtherProps{
		"5'8", 145,
	},
	}
	values, err := Values(p)
	if err != nil {
		t.Errorf("Values(%q) returned error: %v", p, err)
	}
	encoded := values.Encode()
	fmt.Println(encoded)
}

func Test3(t *testing.T) {
	return
	someMap := map[string]interface{}{
		"name":  "Drew",
		"age":   30,
		"beard": false,
	}
	//someSlice := []string{"Hello", "world"}
	values, err := Values(someMap)
	if err != nil {
		t.Errorf("Values(%q) returned error: %v", someMap, err)
	}
	encoded := values.Encode()
	fmt.Println(encoded)
}

func Test4(t *testing.T) {
	someMap := map[string]interface{}{
		"a":  "b",
		"c":  1,
		"d":  false,
		"e":  []interface{}{"f", "g", 2},
		"fo": []string{"g", "h"},
		//"h": map[string]interface{}{
		//	"i": "j",
		//	"k": []interface{}{"l", "m", "n"},
		//},
	}
	//someSlice := []string{"Hello", "world"}

	values, err := Values(someMap)
	if err != nil {
		t.Errorf("Values(%q) returned error: %v", someMap, err)
	}
	encoded := values.Encode()
	unescaped, err := url.QueryUnescape(encoded)
	if err != nil {
		t.Fatal("could not escape")
	}
	fmt.Println(unescaped)
}

func Test5(t *testing.T) {
	return
	someMap := map[string]string{
		"a": "b",
		"c": "d",
	}
	values, err := Values(someMap)
	if err != nil {
		t.Errorf("Values(%q) returned error: %v", someMap, err)
	}
	encoded := values.Encode()
	fmt.Println(encoded)
}
