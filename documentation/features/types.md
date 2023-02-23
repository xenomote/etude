# Structs, Arrays, Maps and Sets

- structs `{...}` are an inhomogeneous type collection of values, with no implied ordering, accessible by name
- arrays `[]T` are a homogeneous type collection of values, with implied ordering, accessible by index
- arrays may optionally have a fixed size, which is a feature of its type i.e. `[10]num8`
- maps `[K]V` are a homogeneous collection of unique keys which map to values, with no implied ordering, accessible through the keys
- sets `[T]` are a homogeneous type collection of unique values, with no implied ordering, accessible only through existence checks
- language keywords to obtain the length of an array, number of keys in a map or the number of elements in a set
- language keywords to apply a sort function based on element values, to turn a set into an array, or to reorder an array
- language keywords to check for existence in a map or set
- operators for concatenation, union, intersection, difference, merge, append etc where appropriate 
- string is just `[]num8[0:]`, no null byte
- `[T]bit` differs from `[T]` because getting a non-existent value from a map creates an error point while a set returns false

# Structural Typing

- types have "simplification" rules which permit a "specific" type to "simplify" into a "simpler" one
- a value of type `A` can be used anywhere a value of type `B` is expected as long as `A` simplifies to `B`
- any type always simplifies to itself as a base case
- a `num` type simplifies to any `num` type containing its entire range with equal or greater precision
- an array `[]A` simplifies to `[A]`, `[]A`
- an array `[n]A` simplifies to `[A]`, `[]A`, `[m]A`, where m <= n
- a map `[K]V` simplifies to `[K]`, `[V]`, `[{K, V}]`
- a set `[A]` simplifies to `[B]` if `A` simplifies to `B`
- a set `[[A]]` simplifies to `[A]`
- a struct `{A}` simplifies to `A`
- a struct `{A, A, A, ...}` simplifies to `[A]`
- a struct simplifies to a struct with fields removed or renamed
- `A` simplifies to `B` if there is a function from `A` to `B`
- named types will guide default behaviour but if names do not provide a solution they can be ignored
- when there are ambiguities in simplification the compiler will present the options and the corresponding clarification to add

# Automatic Type Inference

- function signatures are inferred automatically from the usages of their parameters and the value returned
- parameters are inferred to be as simplified as possible
- return values are inferred to be as specific as possible
- variable declarations without type are inferred to have their most specific usage type
- variable declarations with types are checked to ensure they are at least as specific as their usages
- when inference requires two incompatible specific types then a type error is raised