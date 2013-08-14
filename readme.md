### Gspec
A thin helper around `testing.T` to make writing tests in Go slightly less painful.

### Expect
The main purpose of this library is to streamline the `if false { t.Error(...) }` code which litters Go tests. This is achieved by wrapping the built-in `*T`* class:

    func TestSomething(t *testing.T) {
      spec := gspec.New(t)
      spec.Expect(user.New("goku").PowerLevel).ToEqual(9001)
    }

Methods currently supported are `ToEqual(value)`, `ToBeNil()` and `ToNotBeNill()`

When multiple values are used to set up the expectation, only the first value is compared. This is largely meant to streamline functions which return an error.

### Builers
The library also includes builders for common objects:

  // *http.Request
    req = gspec.Request()
                .Url("/about?home=true")
                .Method("POST")
                .RemoteAddr("119.81.31.35:49958")
                .Header("Accept-Encoding", "gzip").Req
