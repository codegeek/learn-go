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
		{
			"Pointers to things",
			&Person{
				"Guillermo",
				Profile{46, "Monterrey"},
			},
			[]string{"Guillermo", "Monterrey"},
		},
		{
			"Slices",
			[]Profile{
				{46, "Monterrey"},
				{33, "London"},
			},
			[]string{"Monterrey", "London"},
		},
		{
			"Arrays",
			[2]Profile{
				{46, "Monterrey"},
				{33, "London"},
			},
			[]string{"Monterrey", "London"},
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

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
