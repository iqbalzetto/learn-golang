In Go, an array is a numbered sequence of elements of a specific length.
In typical Go code, slices are much more common; arrays are useful in some special scenario

Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part
of the array's type. By default an array is zero-valued, which for ints means 0s.

We can set a value at an index using the array[index] = value syntax, and get a value with array[index].

The builtin len returns the length of an array.

Use this syntax to declare and initialize an array in one line.

Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.