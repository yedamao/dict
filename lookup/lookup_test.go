package lookup

import "testing"
import "fmt"

func TestLookup1(t *testing.T) {
	food := Lookup("go")
	food.PrintAll()
	if food.Word != "go" {
		t.Error("not found the word")
	}
}

func TestLookup2(t *testing.T) {
	food := Lookup("A")
	food.PrintAll()
	if food.Word != "A" {
		t.Error("not found the word")
	}
}

func TestLookup3(t *testing.T) {
	food := Lookup("apple")
	food.PrintAll()
	fmt.Println(food)
	if food.Word != "apple" {
		t.Error("not found the word")
	}
}

func TestLookup4(t *testing.T) {
	food := Lookup("Apple")
	food.PrintAll()
	fmt.Println(food)
	if food.Word != "Apple" {
		t.Error("not found the word")
	}
}

func TestLookup5(t *testing.T) {
	food := Lookup("zoom")
	food.PrintAll()
	fmt.Println(food)
	if food.Word != "zoom" {
		t.Error("not found the word")
	}
}

func TestLookup6(t *testing.T) {
	food := Lookup("*&^%")
	food.PrintAll()
	fmt.Println(food)
	if food.Word != "&^%" {
		t.Error("not found the word")
	}
}

func TestLookup7(t *testing.T) {
	food := Lookup("你")
	food.PrintAll()
	fmt.Println(food)
	if food.Word != "你" {
		t.Error("not found the word")
	}
}
