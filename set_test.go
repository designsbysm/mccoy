package mccoy

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	set := Set{}

	if size := set.Size(); size != 0 {
		t.Errorf("should be empty, got %d", size)
	}

	set.Add("item1")
	if size := set.Size(); size != 1 {
		t.Errorf("should have 1, got %d", size)
	}

	set.Add("item1")
	if size := set.Size(); size != 1 {
		t.Errorf("should have 1, got %d", size)
	}

	set.Add("item2")
	if size := set.Size(); size != 2 {
		t.Errorf("should have 2, got %d", size)
	}
}

func TestClear(t *testing.T) {
	set := Set{}

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")
	set.Clear()

	if size := set.Size(); size != 0 {
		t.Errorf("should be empty, got %d", size)
	}
}

func TestDelete(t *testing.T) {
	set := Set{}

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")
	set.Delete("item2")

	if size := set.Size(); size != 2 {
		t.Errorf("should delete, expected 2 and got %d", size)
	}
}

func TestHas(t *testing.T) {
	set := Set{}

	set.Add("item1")
	set.Add("item2")
	set.Add("item3")
	set.Delete("item2")

	if ok := set.Has("item1"); !ok {
		t.Errorf("should have item1")
	}

	if ok := set.Has("item4"); ok {
		t.Errorf("should not have item1")
	}
}

func TestItems(t *testing.T) {
	set := Set{}

	set.Add("item1")
	set.Add(2)
	set.Add(3.14)
	items := set.Items()

	if len(items) != 3 {
		t.Errorf("should have 3, got %d", len(items))
	}

	if reflect.TypeOf(items[0]).String() != "string" {
		t.Errorf("first item should be string, got %s", reflect.TypeOf(items[0]).String())
	}

	if reflect.TypeOf(items[1]).String() != "int" {
		t.Errorf("first item should be int, got %s", reflect.TypeOf(items[1]).String())
	}

	if reflect.TypeOf(items[2]).String() != "float64" {
		t.Errorf("first item should be float, got %s", reflect.TypeOf(items[2]).String())
	}
}

func TestSzie(t *testing.T) {
	set := Set{}

	set.Add("item1")
	set.Add(2)
	set.Add(3.14)
	items := set.Items()

	if len(items) != 3 {
		t.Errorf("should have 3, got %d", len(items))
	}
}
