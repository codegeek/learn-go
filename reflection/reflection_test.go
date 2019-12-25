package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Guillermo"},
			[]string{"Guillermo"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Guillermo", "Monterrey"},
			[]string{"Guillermo", "Monterrey"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Guillermo", 46},
			[]string{"Guillermo"},
		},
		{
			"Nested fields",
			Person{
				"Guillermo",
				Profile{46, "Monterrey"},
			},
			[]string{"Guillermo", "Monterrey"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
