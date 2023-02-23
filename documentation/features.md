# Arbitrary Byte Width Numbers

- `num` keyword suffixed with a 8-multiple bit width, i.e. `num8`, `num32` etc
- specifying either min or max offsets the range to have that min or max, i.e. `num8[0:]`, `num64[:0]`
- specifying both min and max replaces bit width, and limits range, i.e `num[6:10]`
- specifying power of 10 precision allows decimal values and floating point arithmetic, i.e. `num32[-3]`
- both range and precision may be specified simultaneously, i.e `num[-2E20:2E20,+4]`
- redundant number types are not permitted, i.e. `num[x:x,y]`, `num[x:x+1,y]`, `num[x2:x1,z]`, `num[x1:x2,0]` for any x, x1 < x2 and y
- implemented by the most appropriate hardware type, range translations and software arithmetic used if necessary
- language keywords to obtain the max and min value of a type as a constant
- overflows, underflows, truncation, etc are all statically checked and create error points if necessary
- `bit` is the type for boolean operations, backed by the most appropriate hardware type, may be coalesced with other `bit`s

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

# No Global State

- only functions and constants are accessible at global scope
- standard library designed so that kernel functionality is accessed statelessly

# Separate Files

- error handling, logging and documentation concerns are handled in separate files to source code
- shares some common aspects of the language, but uses specific syntax where necessary
- compiler/editor support which cross-references these as appropriate
- compiler has a method of checking for dependencies on the source file and requires an update

# Error Handling

- checked and enforced statically by the compiler to create error "points", errors cannot be ignored
- keywords for different error handling strategies to apply at any given point
- termination policy is permitted but only from the main function
- retry policies can specify number of attempts, timeouts, conditions
- adjustment policy can modify a value in scope at the point of error
- different strategies may annotate the signature of the function
- error handling gets access to the values in scope at the point of execution

# Logging

- define messages to log at specific points of execution
- use variables from the scope in messages or conditions
- add, remove or update logging statements at runtime

# Documentation

- checked and enforced by the compiler
- can be placed as inline references on any language structure or grouping
- can call arbitrary tools to retrieve specific text
- generates warnings when documentation is updated, shows diff, requires review
- can be used to splice text into the source file

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

# Automatic Concurrency and Asynchronicity

- arrays may be treated as buffered channels at runtime and used to allow parallelism
- hardware/driver/io calls in one statement may cause a context switch to other statements
- repeated timer interrupts in one statement may cause statements to be scheduled to another core if available
- no explicit locking or concurrency primitives in the language, all handled by the runtime
- creates config files on compilation which are used to tune specific concurrency parameters

# Statement Grouping

- a sequence of adjacent statements is a group structure
- group structures are separated from other group structures by blank lines
- may consist of a linear sequence of data dependent statements (array)
- may consist of a group of data independent statements which reference a cohesive set of variables (set)
- groups may also be nested in larger groups based on cohesion, although this is implied through variable usage
- full grouping structure can be displayed by the compiler to show semantic interpretation
- used to guide concurrency and asynchronicity, although may also be restructured to serve it

# Automatic Code Rearrangement

- static analysis to optimise cohesion within functions, files and packages
- reorganisation of function flows in order to create sensible groupings and orderings
- aggressive detection and automatic replacement of repeated structure
- aggressive decomposition of complex functionality into subfunctions

# Mini Executables

- unified system for directly calling functions from the command line
- automatic generation of usage, help, etc
- automatic conversion from common command line formats, e.x. pipe, file, string literal
- file for defining default arguments, config options, argument for overriding configs