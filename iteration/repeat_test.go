package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat the character 5 times",
		func(t *testing.T) {
			repeated := Repeat("a", 10)
			expected := "aaaaaaaaaa"

			if repeated != expected {
				t.Errorf("expected %q, got %q", expected, repeated)
			}
		})

	t.Run("Repeat the character 5 times when count is negative",
		func(t *testing.T) {
			repeated := Repeat("a", -1)
			expected := "aaaaa"

			if repeated != expected {
				t.Errorf("expected %q, got %q", expected, repeated)
			}
		})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
