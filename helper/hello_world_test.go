package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Ahmad",
			request: "Ahmad",
		},
		{
			name:    "Muhbit",
			request: "Muhbit",
		},
		{
			name:    "AhmadMuhbit",
			request: "Ahmad Muhbit",
		},
		{
			name:    "MuhammadAltafAdaby",
			request: "Muhammad Altaf Adaby",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Ahmad", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ahmad")
		}
	})
	b.Run("Muhbit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Muhbit")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ahmad")
	}
}

func BenchmarkHelloWorldMuhbit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Muhbit")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ahmad",
			request:  "Ahmad",
			expected: "Hello Ahmad",
		},
		{
			name:     "Muhbit",
			request:  "Muhbit",
			expected: "Hello Muhbit",
		},
		{
			name:     "Altaf",
			request:  "Altaf",
			expected: "Hello Altaf",
		},
		{
			name:     "Lina",
			request:  "Lina",
			expected: "Hello Lina",
		},
		{
			name:     "Nadhira",
			request:  "Nadhira",
			expected: "Hello Nadhira",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Ahmad", func(t *testing.T) {
		result := HelloWorld("Ahmad")
		require.Equal(t, "Hello Ahmad", result, "Result must be 'Hello Ahmad'")
	})

	t.Run("Muhbit", func(t *testing.T) {
		result := HelloWorld("Muhbit")
		require.Equal(t, "Hello Muhbit", result, "Result must be 'Hello Muhbit'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before unit test")
	m.Run()

	// after
	fmt.Println("After unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Ahmad")
	require.Equal(t, "Hello Ahmad", result, "Result must be 'Hello Ahmad'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ahmad")
	require.Equal(t, "Hello Ahmad", result, "Result must be 'Hello Ahmad'")
	fmt.Println("TestHelloWorldAhmad with Require done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ahmad")
	assert.Equal(t, "Hello Ahmad", result, "Result must be 'Hello Ahmad'")
	fmt.Println("TestHelloWorldAhmad with Assert done")
}

func TestHelloWorldAhmad(t *testing.T) {
	result := HelloWorld("Ahmad")

	if result != "Hello Ahmad" {
		// error
		//t.Fail()
		t.Error("Result must be 'Hello Ahmad'") // disarankan, karena bisa melihar errornya apa
	}

	fmt.Println("TestHelloWorldAhmad Done")
}

func TestHelloWorldMuhbit(t *testing.T) {
	result := HelloWorld("Muhbit")

	if result != "Hello Muhbit" {
		// error
		//t.FailNow()
		t.Fatal("Result must be 'Hello Muhbit'") // disarankan, karena bisa melihar errornya apa
	}

	fmt.Println("TestHelloWorldMuhbit Done")
}
