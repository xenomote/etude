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

`.` takes all names from preceeding expression and makes
them availabe to the scope of the following expression

`,` does the same as `.` but merges the following expression onto the preceeding expression. must be two scopes, two tuples or two lists

```
{
    min [a; ...bs (b; ...rest)] num:
        (bs.empty)? a,
        [(a < b)? a, b; rest...].min;
    
    sub [a; ...bs (b; ...rest)] num:
        (bs.empty)? a,
        [(a - b); rest...].sub;
}.{
    gcd (ab [a; b]) num: 
        (ab.equal)? a,
        [ab.sub; ab.min].gcd;
}
```

```
{

}
```