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