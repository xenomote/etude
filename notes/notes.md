should numbers be a tuple of bits?
i.e. 
- `num := [bit]`
- `byte := [bit 8]`
- `decimal := {coefficient: num; power: num}`
- `float := {mantissa: num; exponent: num}`

function bodies can be sequences of object merges
i.e.
```
{
    distance (x; y) num: (x.squared; y.squared).sum.square root;
},{
    to50 num num: @50.distance;
}.(
    (44).to50
)
```

`.` takes all names from preceeding expression and makes them availabe to the scope of the following expression

`,` does the same as `.` but merges the following expression onto the preceeding expression. must be two scopes, two tuples or two lists

`?` does the same as `.` but makes the following expression optional

```
{
    gcd (a; b) ~: a = b? a, gcd(a - b; (a < b)? a, b);
}.(
    gcd (23, 44)
)
```

`...a` is gather, `a...` is spread


need to express homogeny vs inhomogeny, ordered vs unordered

homogenous ordered is list, `[T]`, literal `[1, 2, 3, x..., 4]`
homogenous unordered is set, `{~T}`, literal `{}`
inhomogenous ordered is record, `(T, U, V)`
inhomogenous unordered is scope, `{a T, b U, c V}`, requires names


```
<num> := <num> <bit?> := <list bit>

3(0)?none = true
3(7)?none = none

<list t> := <num> <t?>

[1, 2, 3, 4](0)?none = 1
[1, 2, 3, 4](5)?none = none

<set t> := <t> <bit>

{1, 3, 5}(1) = true
{1, 3, 5}(2) = false

<map {key k, value v}> := <k> <v?>

{1: a, 3: b, 5: c}(3)?none = b
{1: a, 3: b, 5: c}(2)?none = none

<t?> ? <u> -> t? <t|u>
<t|t> -> <t>

<map num>?<num> -> <num>

```