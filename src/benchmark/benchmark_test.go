package benchmark

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/fulldump/apitest"

	"chi_benchmark"
	"golax_benchmark"
	"gorilla_benchmark"
)

func GetImplementations() map[string]http.Handler {
	return map[string]http.Handler{
		"chi":     chi_benchmark.NewApi(),
		"golax":   golax_benchmark.NewApi(),
		"gorilla": gorilla_benchmark.NewApi(),
	}
}

func ListUsers(b *testing.B, a *apitest.Apitest) {
	a.Request("GET", "/service/v1/users").DoAsync(func(r *apitest.Response) {
		r.BodyString()
	})
}

func GetUser(b *testing.B, a *apitest.Apitest) {
	a.Request("GET", "/service/v1/users/2").DoAsync(func(r *apitest.Response) {
		r.BodyString()
	})
}

func GetZZ(b *testing.B, a *apitest.Apitest) {
	a.Request("GET", "/letters/z/z").DoAsync(func(r *apitest.Response) {
		r.BodyString()
	})
}

func GetZZZ(b *testing.B, a *apitest.Apitest) {
	a.Request("GET", "/letters/z/z/z").DoAsync(func(r *apitest.Response) {
		r.BodyString()
	})
}

type work func(*testing.B, *apitest.Apitest)

func RunParallel(b *testing.B, a *apitest.Apitest, f work, p int) {

	c := make(chan bool, p*4)
	go func() {
		for i := 0; i < b.N; i++ {
			c <- false
		}
		for i := 0; i < p; i++ {
			c <- true
		}
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < p; i++ {
		wg.Add(1)
		go func() {
			for {
				stop := <-c
				if stop {
					wg.Done()
					return
				}

				f(b, a)
			}
		}()
	}
	wg.Wait()

}

func BenchmarkPerformance(b *testing.B) {

	b.ReportAllocs()
	b.SetBytes(1)

	benchmarks := map[string]work{
		"list users":           ListUsers,
		"retrieve user":        GetUser,
		"retrieve letters zz":  GetZZ,
		"retrieve letters zzz": GetZZZ,
	}

	threads := []int{1, 10, 100}

	for _, t := range threads {
		fmt.Println("\n#", t, "threads")

		for framework, implementation := range GetImplementations() {
			fmt.Println("\n##", framework)

			a := apitest.NewWithPool(implementation, t)
			for benchmark_name, benchmark := range benchmarks {
				fmt.Println("###", benchmark_name)

				n := 0

				t0 := time.Now()
				b.Run(benchmark_name, func(b *testing.B) {
					n += b.N

					RunParallel(b, a, benchmark, t)
				})
				t1 := time.Now()

				fmt.Println("TPS:", float64(n)/t1.Sub(t0).Seconds())

			}
			a.Destroy()

		}
	}

}
