package filter

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var target []string
var targetInt []int

func startsWithVowel(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func buildString() string {
	out := make([]byte, 5)
	for i := 0; i < 5; i++ {
		out[i] = byte(rand.Intn(26) + 65)
		return string(out)
	}
	return string(out)
}

func setup() []string {
	rand.Seed(1)
	vals := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = buildString()
	}
	return vals
}

func setupInt() []int {
	rand.Seed(2)
	vals := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = i
	}
	return vals
}

func isEven(i int) bool {
	return i%2 == 0
}

func BenchmarkFilterReflectString(b *testing.B) {
	vals := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = ReflectedFilter(vals, startsWithVowel).([]string)
	}
}

func BenchmarkFilterString(b *testing.B) {
	vals := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = Filter(vals, startsWithVowel)
	}
}

func BenchmarkFilterReflectInt(b *testing.B) {
	vals := setupInt()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = ReflectedFilter(vals, isEven).([]int)
	}
}

func BenchmarkFilterInt(b *testing.B) {
	vals := setupInt()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = FilterInt(vals, isEven)
	}
}

func TestFilterReflect(t *testing.T) {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	longNames := ReflectedFilter(names, func(s string) bool {
		return len(s) > 3
	}).([]string)

	fmt.Println((longNames))
	if diff := cmp.Diff(longNames, []string{"Andrew", "Clara", "Hortense"}); diff != "" {
		t.Error(diff)
	}

	ages := []int{20, 50, 13}
	adults := ReflectedFilter(ages, func(age int) bool {
		return age >= 18
	}).([]int)

	fmt.Println(adults)
	if diff := cmp.Diff(adults, []int{20, 50}); diff != "" {
		t.Error(diff)
	}
}
