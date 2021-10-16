package mccoy

import "sort"

type Item interface{}

func ItemString(i Item) string {
	return i.(string)
}

func ItemsSortStrings(items []Item) []string {
	values := []string{}

	for _, i := range items {
		str := ItemString(i)
		values = append(values, str)
	}

	sort.Strings(values)

	return values
}
