package cmp_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	testcmp "github.com/junhyuk0801/learning-go/13-tests/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := testcmp.Person{
		Name: "Dennis",
		Age:  37,
	}

	result := testcmp.CreatePerson("Dennis", 37)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestCreatePerson_IgnoreDate(t *testing.T) {
	expected := testcmp.Person{
		Name: "Dennis",
		Age:  37,
	}
	result := testcmp.CreatePerson("Dennis", 37)
	comparer := cmp.Comparer(func(x, y testcmp.Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
	if result.DateAdded.IsZero() {
		t.Error("DateAdded was not assigned")
	}

}
