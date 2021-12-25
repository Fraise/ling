# ling

ling is a functional programming style library which aims to simplify data processing in Go. It requires Go 1.18 at
minimum as is uses generics to allow manipulating any data.

## Chainable methods

Some methods can be chained together through the use of the `SliceChainer`. This type is a wrapper of `[]T` and allows
methods to be defined for slices.

Methods that cannot be chained are usually those that convert one generic type to another (such as `Map`) or those that
requires additional constraints to work (such as `Min` or `Max`).

The easiest way to use it is with the functions `IntoChainer(in []T)` and `(chainer SliceChainer) ToSlice()`.
