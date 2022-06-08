package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestLen(t *testing.T) {
	t.Parallel()
	tData := map[string]struct {
		A        People
		Expected int
	}{
		"success": {make([]Person, 1), 1},
		"fail":    {make([]Person, 0), 0},
	}

	for name, tcase := range tData {
		v := tcase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := v.A.Len()
			if got != v.Expected {
				t.Errorf("[%s] expected: %d, got %d", name, v.Expected, got)
			}
		})
	}
}

func TestLess(t *testing.T) {
	t.Parallel()
	tData := map[string]struct {
		A        People
		i        int
		j        int
		Expected bool
	}{
		"caseBirthday": {People{Person{birthDay: time.Date(1987, 11, 15, 0, 0, 0, 0, time.UTC)},
			Person{birthDay: time.Date(1987, 05, 18, 0, 0, 0, 0, time.UTC)}}, 0, 1, true},
		"caseFirstNames": {People{Person{birthDay: time.Date(1987, 11, 15, 0, 0, 0, 0, time.UTC), firstName: "Dmitry"},
			Person{birthDay: time.Date(1987, 11, 15, 0, 0, 0, 0, time.UTC), firstName: "Vitaly"}}, 0, 1, true},
		"caseLastNames": {People{Person{birthDay: time.Date(1987, 11, 15, 0, 0, 0, 0, time.UTC), firstName: "Dmitry", lastName: "Tolkachev"},
			Person{birthDay: time.Date(1987, 11, 15, 0, 0, 0, 0, time.UTC), firstName: "Dmitry", lastName: "Donikov"}}, 0, 1, false},
	}

	for name, tcase := range tData {
		v := tcase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := v.A.Less(v.i, v.j)
			if got != v.Expected {
				t.Errorf("[%s] expected: %d, got %d", name, v.Expected, v.A)
			}
		})
	}
}
func TestSwap(t *testing.T) {
	t.Parallel()
	tData := map[string]struct {
		A People
		i int
		j int
	}{
		"swapping": {People{Person{firstName: "Vitaly", lastName: "Donikov", birthDay: time.Date(1987, 05, 18, 0, 0, 0, 0, time.UTC)},
			Person{firstName: "Dmitry", lastName: "Tolkachev", birthDay: time.Date(1987, 11, 15, 0, 0, 0, 0, time.UTC)}}, 0, 1},
	}

	for name, tcase := range tData {
		v := tcase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			v.A.Swap(v.i, v.j)
			if v.A[0].firstName != "Dmitry" || v.A[1].firstName != "Vitaly" {
				if v.A[0].lastName != "Tolkachev" || v.A[1].lastName != "Donikov" {
					if v.A[0].birthDay != time.Date(1987, 11, 15, 0, 0, 0, 0, time.UTC) || v.A[1].birthDay != time.Date(1987, 05, 18, 0, 0, 0, 0, time.UTC) {

						t.Errorf("[%s] expected: %d, got %d", name, v.A[0], v.A[1])
					}
				}
			}
		})

	}
}
func TestNew(t *testing.T) {
	var ErrSize = errors.New("Rows need to be the same length")
	var _, ErrAtoi = strconv.Atoi("a")
	t.Parallel()
	tData := map[string]struct {
		str      string
		Expected *Matrix
		Err      error
	}{
		"succes":              {str: "1 2 3 \n 4 5 6 \n7 8 9", Expected: &Matrix{rows: 3, cols: 3, data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, Err: nil},
		"matrix_size":         {str: "10 11\n 12 13 ", Expected: nil, Err: ErrSize},
		"letters_into_matrix": {str: "a b c\n a x z", Expected: nil, Err: ErrAtoi},
	}

	for name, tCase := range tData {
		v := tCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := New(v.str)
			c := false
			if err != nil && err.Error() != v.Err.Error() {
				t.Errorf("[%s] expected error: %v, got error: %v", name, v.Err, err)
			}
			if v.Expected != nil {
				for k, i := range got.data {
					if v.Expected.data[k] != i {
						c = true
					}
				}
				if got.cols != v.Expected.cols || got.rows != v.Expected.rows || c {
					t.Errorf("[%s] expected: %v, got: %v", name, v.Expected, got)

				}
			}

		})
	}
}
