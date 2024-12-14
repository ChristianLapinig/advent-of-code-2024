package src

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup(m)
	code := m.Run()
	teardown(m)
	os.Exit(code)
}

func teardown(m *testing.M) {
	os.RemoveAll("test")
}

func setup(m *testing.M) {
	os.MkdirAll("test", os.ModePerm)
	file, err := os.Create("test/test.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := []byte("3   4\n4   3\n2   5\n1   3\n3   9\n3   3\n")
	if _, err := file.Write(data); err != nil {
		panic(err)
	}
}

func TestDayOne(t *testing.T) {
	assertEquals := func(t testing.TB, actual, expected int) {
		if actual != expected {
			t.Fatalf("FAILED: got %d, expected %d", actual, expected)
		}
	}

	solution := DayOne{FilePath: "./test/test.txt"}

	t.Run("test CalculateTotalDistance", func(t *testing.T) {
		assertEquals(t, solution.CalculateTotalDistance(), 11)
	})

	t.Run("test CalculateTotalSimilarityScore", func(t *testing.T) {
		assertEquals(t, solution.CalculateTotalSimilarityScore(), 31)
	})
}
