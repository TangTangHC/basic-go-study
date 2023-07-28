package basic

import "testing"

func Test_1(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func Test_2(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}

func Test_3(t *testing.T) {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
