package set

import (
	"reflect"
	"testing"
)

func TestItems(t *testing.T) {
	set := New()

	set.Add("item1")
	set.Add(2)
	set.Add(3.14)
	items := set.Items()

	if len(items) != 3 {
		t.Errorf("should have 3, got %d", len(items))
	}

	var stringFound bool
	var intFound bool
	var floatFound bool

	for _, value := range items {
		if value == "item1" && reflect.TypeOf(value).String() == "string" {
			stringFound = true
		}
		if value == 2 && reflect.TypeOf(value).String() == "int" {
			intFound = true
		}
		if value == 3.14 && reflect.TypeOf(value).String() == "float64" {
			floatFound = true
		}
	}

	if !stringFound {
		t.Errorf("items should contain string, intem1")
	}

	if !intFound {
		t.Errorf("items should contain int, 2")
	}

	if !floatFound {
		t.Errorf("items should contain float64, 3.14")
	}
}
