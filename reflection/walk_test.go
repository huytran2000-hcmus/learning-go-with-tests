package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}

	simpleCases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string",
			struct {
				Name string
			}{"Huy"},
			[]string{"Huy"},
		},
		{
			"struct with two strings",
			struct {
				Name string
				City string
			}{"Huy", "HCM"},
			[]string{"Huy", "HCM"},
		},
		{
			"struct with non string",
			struct {
				Name string
				Age  int
			}{"Huy", 22},
			[]string{"Huy"},
		},
		{
			"struct with nested struct",
			Person{
				"Huy",
				Profile{
					22,
					"HCM",
				},
			},
			[]string{"Huy", "HCM"},
		},
		{
			"pointer to struct",
			&Person{
				"Huy",
				Profile{
					22,
					"HCM",
				},
			},
			[]string{"Huy", "HCM"},
		},
		{
			"slice",
			[]Profile{
				{22, "Huy"},
				{22, "Kha"},
			},
			[]string{"Huy", "Kha"},
		},
		{
			"array",
			[2]Profile{
				{22, "Huy"},
				{22, "Kha"},
			},
			[]string{"Huy", "Kha"},
		},
	}

	for _, c := range simpleCases {
		t.Run(c.Name, func(t *testing.T) {
			var got []string
			Walk(c.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, c.ExpectedCalls) {
				t.Errorf("want %q, got %q", c.ExpectedCalls, got)
			}
		})
	}

	t.Run("map", func(t *testing.T) {
		input := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		want := []string{"Bar", "Boz"}
		var got []string

		Walk(input, func(result string) {
			got = append(got, result)
		})

		if len(want) != len(got) {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), len(want))
		}

		for _, item := range got {
			assertContains(t, want, item)
		}
	})
	t.Run("channel", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{22, "Huy"}
			ch <- Profile{22, "Khai"}
			close(ch)
		}()

		want := []string{"Huy", "Khai"}
		var got []string

		Walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{22, "Huy"}, Profile{22, "Kha"}
		}

		want := []string{"Huy", "Kha"}
		var got []string
		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertContains(t testing.TB, wants []string, got string) {
	t.Helper()

	contains := false
	for _, x := range wants {
		if got == x {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %q to contain %q but it didn't", wants, got)
	}
}
