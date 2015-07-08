#ID Mask
-------------

Convert a slice of unsigned integers into an integer bitmask and back again.

## Known issues
-------------

Unable to handle numbers larger than 62. See failing test.

```
$ go test
--- FAIL: TestIDMaskLarger (0.00 seconds)
    mask_test.go:32: got [1 2 3 5 6 8 62], want [1 3 5 8 140 150 6 2 62 63 63 127 128 129]
FAIL
exit status 1
```