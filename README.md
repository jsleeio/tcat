# tcat

## what?

`tcat` is much like `cat` but transparently decompresses files named with `.gz`
suffix as `gzip`, and files named with `.bz2` suffix as `bzip2`. Other files
are treated as uncompressed.

## options

```
$ tcat -help
Usage of ./tcat:
  -summary
    	summarize file types and line counts only
```

## license

Copyright 2019 John Slee.  Released under the terms of the MIT license
[as included in this repository](LICENSE).
