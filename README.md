zipmap provides utility functions for reading and writing a limited
range of Go maps as gzipped files and streams, using gob as the wire
format.  Currently, a limited set of map types with primitive keys and
values is supported: `map[int64]float64`, `map[int64]float64`, and
`map[int64]string`.

Example of writing a map named `mp` (note that this can actually write
any object type supported by gob):

```
err := zipmap.Write(mp, "file.bin.gz")
if err != nil {
 	  panic(err)
}
```

Example of reading a `map[int64]float64`

```
mp, err := zipmap.ReadInt64Float64("file.bin.gz")
if err != nil {
 	  panic(err)
}
```
