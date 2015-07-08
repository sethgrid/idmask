# ID Mask
-------------

Convert a slice of unsigned integers into an integer bitmask and back again.

### Technique
-------------

Because we cannot store a bitmask of a number larger than 64 in a bitmask, the mask utility creates buckets (elements in a slice) that represent numbers 0-63, 64-127, 127-....

### Example
-------------
Given the following list of integers:

`[1 3 5 8 140 150 6 72 62 63 64 65 127 128 129 130]`

We can generate the following buckets that represent the bitmasks:

`[13835058055282164074 9223372036854776067 4198407]`

The element at index 0 represents numbers below 64, the element and index 1 represents those below 128, and the element in index 3 represents numbers below 192.
