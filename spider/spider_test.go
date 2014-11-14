package spider

import "testing"

func TestSpider0(t *testing.T) {
    food := Spider("apple")
    food.PrintAll()
    food.WriteAll("./test.dict")
    if food.Word != "apple" {
        t.Error("not the find word")
    }
}

func TestSpider1(t *testing.T) {
    food := Spider("go")
    food.PrintAll()
    food.WriteAll("./test.dict")
    if food.Word != "go" {
        t.Error("not the find word")
    }
}

func TestSpider2(t *testing.T) {
    food := Spider("hhh")
    food.PrintAll()
    food.WriteAll("./test.dict")
    if food.Word != "hhh" {
        t.Error("not the find word")
    }
}

func TestSpider3(t *testing.T) {
    food := Spider("hasdxxx")
    food.PrintAll()
    food.WriteAll("./test.dict")
    if food.Word != "" {
        t.Error("not the find word")
    }
}

func TestSpider4(t *testing.T) {
    food := Spider("好")
    food.PrintAll()
    food.WriteAll("./test.dict")
    if food.Word != "" {
        t.Error("not the find word")
    }
}

func TestSpider5(t *testing.T) {
    food := Spider("good")
    food.PrintAll()
    food.WriteAll("./test.dict")
    if food.Meaning != "good" {
        t.Error("not the find word")
    }
}
