A simple mapper function to convert between domain objects such as REST entity and a service layer. 
The Mapper function copies the values for exact field name matches between the source and destination structure.

Example:
```
type Source struct {
	SourceOnly string
	Common string
}

type Destination struct {
	DestinationOnly int
	Common string
}

s := &Source{"SourceOnly", "Common"}
d := &Destination{DestinationOnly:1}
Mapper(s, d)
```
Mapper copies the value of the field `Common`. 
```
fmt.Println(d)
&{1 Common}
```
