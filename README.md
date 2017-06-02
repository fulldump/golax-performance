# golax performance

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [Intro](#intro)
- [The results](#the-results)
	- [List users](#list-users)
	- [Retrieve user](#retrieve-user)
	- [`GET letters/z/z`](#get-letterszz)
	- [`GET letters/z/z/z`](#get-letterszzz)
	- [Keep alive with 100 threads](#keep-alive-with-100-threads)
- [Run tests in your machine](#run-tests-in-your-machine)
- [About the implementations](#about-the-implementations)
	- [Golax](#golax)
	- [Gorilla](#gorilla)
	- [Chi](#chi)
- [About the code readability and maintainability](#about-the-code-readability-and-maintainability)

<!-- /MarkdownTOC -->

## Intro

The reason for this project is [this question](https://groups.google.com/forum/#!msg/golang-nuts/W8oETGFBu_o/Z4glNpoiGgAJ) made by [Adrian Sampaleanu](https://plus.google.com/+AdrianSampaleanu) in _golang-nuts_.

The performance compared with the [most popular alternative](http://www.gorillatoolkit.org/) is very similar (actually _golax_ performs slightly better) however code readability and maintainability is far better with _golax_ implementation.

## The results

### List users

<p align="center">
    <img src="https://docs.google.com/spreadsheets/d/1q0NdoBge4UO_VmFGwcYDN4WQZzKuqCXNrtJzThVdJWQ/pubchart?oid=1063158416&format=image">
</p>

### Retrieve user

<p align="center">
    <img src="https://docs.google.com/spreadsheets/d/1q0NdoBge4UO_VmFGwcYDN4WQZzKuqCXNrtJzThVdJWQ/pubchart?oid=478921787&format=image">
</p>

### `GET letters/z/z`

<p align="center">
    <img src="https://docs.google.com/spreadsheets/d/1q0NdoBge4UO_VmFGwcYDN4WQZzKuqCXNrtJzThVdJWQ/pubchart?oid=1350397502&format=image">
</p>

### `GET letters/z/z/z`

<p align="center">
    <img src="https://docs.google.com/spreadsheets/d/1q0NdoBge4UO_VmFGwcYDN4WQZzKuqCXNrtJzThVdJWQ/pubchart?oid=1153847898&format=image">
</p>

### Keep alive with 100 threads

<p align="center">
    <img src="https://docs.google.com/spreadsheets/d/1q0NdoBge4UO_VmFGwcYDN4WQZzKuqCXNrtJzThVdJWQ/pubchart?oid=1936169051&format=image">
</p>


Tests has been executed in a `Intel(R) Core(TM) i5-2400 CPU @ 3.10GHz` and 16GiB RAM.

## Run tests in your machine

Run all benchmarks for all frameworks:

```sh
make dependencies
make benchmark
```

Make and run golax:

```sh
make golax
```

Make and run gorilla:
```sh
make gorilla
```

Make and run chi:
```sh
make chi
```

Execute tests:

1 thread:
```sh
ab -n 20000 -c 1 http://localhost:9999/service/v1/users
ab -n 20000 -c 1 http://localhost:9999/service/v1/users/2
ab -n 20000 -c 1 http://localhost:9999/letters/z/z
ab -n 20000 -c 1 http://localhost:9999/letters/z/z/z
```

10 threads:
```sh
ab -n 200000 -c 10 http://localhost:9999/service/v1/users
ab -n 200000 -c 10 http://localhost:9999/service/v1/users/2
ab -n 200000 -c 10 http://localhost:9999/letters/z/z
ab -n 200000 -c 10 http://localhost:9999/letters/z/z/z
```

100 threads:
```sh
ab -n 200000 -c 100 http://localhost:9999/service/v1/users
ab -n 200000 -c 100 http://localhost:9999/service/v1/users/2
ab -n 200000 -c 100 http://localhost:9999/letters/z/z
ab -n 200000 -c 100 http://localhost:9999/letters/z/z/z
```

500 threads:
```sh
ab -n 200000 -c 500 http://localhost:9999/service/v1/users
ab -n 200000 -c 500 http://localhost:9999/service/v1/users/2
ab -n 200000 -c 500 http://localhost:9999/letters/z/z
ab -n 200000 -c 500 http://localhost:9999/letters/z/z/z
```

Keep alive:
```sh
ab -k -n 20000 -c 100 http://localhost:9999/service/v1/users
ab -k -n 20000 -c 100 http://localhost:9999/service/v1/users/2
ab -k -n 20000 -c 100 http://localhost:9999/letters/z/z
ab -k -n 20000 -c 100 http://localhost:9999/letters/z/z/z
```


## About the implementations

All implement a CRUD API described [here](https://github.com/fulldump/golax/blob/master/example/README.md) and:

* Errors are returned in JSON format
* All requests are logged to stdout and/or stderr


### Golax

The code is the standard way a REST API should be implemented with golax.


### Gorilla

Gorilla implementation has been done following [Making a RESTful JSON API in Go](https://thenewstack.io/make-a-restful-json-api-go/) article.


### Chi

I am glad to know about Chi, it follows the same approach as golax and it has a very similar implementation. 


## About the code readability and maintainability

TODO: comment several points here, and how easy (or not) is adding middlewares/interceptors and routes.
