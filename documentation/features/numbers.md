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