package mccoy

import (
	"reflect"
	"testing"
)

func TestItemString(t *testing.T) {
	set := Set{}

	set.Add("item1")
	str := ItemString(set.Items()[0])

	if str != "item1" {
		t.Errorf("value should equal item1, got %s", str)
	}

	if reflect.TypeOf(str).String() != "string" {
		t.Errorf("item should be string, got %s", reflect.TypeOf(str).String())
	}
}

func TestItemsSortStrings(t *testing.T) {
	set := Set{}

	set.Add("c")
	set.Add("b")
	set.Add("a")
	items := ItemsSortStrings(set.Items())

	if items[0] != "a" || items[1] != "b" || items[2] != "c" {
		t.Errorf("items should be sorted [a b c], got %s", items)
	}
}
