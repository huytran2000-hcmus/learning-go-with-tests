package generics

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSum(t *testing.T) {
	t.Run("sum of arbitrary size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		AssertEqual(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 1, 9})
	want := []int{6, 10}

	AssertEqual(t, got, want)
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slice",
		func(t *testing.T) {
			got := SumAllTails([]int{1, 2}, []int{0, 9})
			want := []int{2, 9}

			AssertEqual(t, got, want)
		})

	t.Run("safely sum empty slices",
		func(t *testing.T) {
			got := SumAllTails([]int{}, []int{})
			want := []int{0, 0}

			AssertEqual(t, got, want)
		})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(accumulated, val int) int {
			return accumulated * val
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(accumalated, val string) string {
			return accumalated + val
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func AssertEqual[T any](t testing.TB, got, want T) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("got %v, want %v", got, want)
	}
}
