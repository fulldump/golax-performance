package benchmark

import (
	"fmt"
	"net/http"
	"testing"

	"gorilla_benchmark"

	"chi_benchmark"
	"golax_benchmark"

	"github.com/fulldump/apitest"
)

func GetImplementations() map[string]http.Handler {
	return map[string]http.Handler{
		"chi":     chi_benchmark.NewApi(),
		"golax":   golax_benchmark.NewApi(),
		"gorilla": gorilla_benchmark.NewApi(),
	}
}

func ListUsers(b *testing.B, w int, t *apitest.Apitest) {
	t.Request("GET", "/service/v1/users").DoParallel(b.N, w, func(r *apitest.Response, err error) {
		//fmt.Println(r.StatusCode)
	})
}

func GetUser(b *testing.B, w int, t *apitest.Apitest) {
	t.Request("GET", "/service/v1/users/2").DoParallel(b.N, w, func(r *apitest.Response, err error) {
		//fmt.Println(r.StatusCode)
	})
}

func GetZZ(b *testing.B, w int, t *apitest.Apitest) {
	t.Request("GET", "/letters/z/z").DoParallel(b.N, w, func(r *apitest.Response, err error) {
		//fmt.Println(r.StatusCode)
	})
}

func GetZZZ(b *testing.B, w int, t *apitest.Apitest) {
	t.Request("GET", "/letters/z/z/z").DoParallel(b.N, w, func(r *apitest.Response, err error) {
		//fmt.Println(r.StatusCode)
	})
}

type PrepareRequest func(*testing.B, int, *apitest.Apitest)

func BenchmarkPerformance(b *testing.B) {

	b.ReportAllocs()
	b.SetBytes(2)

	benchmarks := map[string]PrepareRequest{
		"list users":           ListUsers,
		"retrieve user":        GetUser,
		"retrieve letters zz":  GetZZ,
		"retrieve letters zzz": GetZZZ,
	}

	for framework, implementation := range GetImplementations() {
		fmt.Println("#", framework)

		a := apitest.NewApitest(implementation)
		for benchmark_name, benchmark := range benchmarks {
			fmt.Println("##", benchmark_name)
			b.Run(benchmark_name, func(b *testing.B) {
				benchmark(b, 10, a)
			})
		}

	}
}
