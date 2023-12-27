# redundantptr

The redundantptr tool is a static analisys tool for golang code. It finds assignments of pointers to variables.

For example:
```golang
var value = "some value"
s := myType{aPointerField: &value}
```

This is usually done because go does not support pointers to constant values. Here is another common example:

```golang
const someConstant = int32(42)

...

value := someConstant
s := myType{int32PointerField: &value}
```

In projects that uses kubernetes packages, it is possible to use the `ptr.To()` function to get rid of these variables. For example:
```
```golang
const someConstant = int32(42)

...

s := myType{int32PointerField: ptr.To(someConstant)}
```

The redundantptr tool help to find these cases. 
### Limitations
The tool is not perfect. I currently will also notify about cases that do make sense to use variables.
Use the tool carefully and always check the result manually.

For example, in the following code, we need to have some logic aroud the desired value to be assigned. In this case we do want to use a variable:
```golang
const (
  a = 5
  b = 7
)

...

var value := a
if someCondition() {
  value = b
}

s := myType{aPointerField: &value}
```

## Build
To build the tool, run:
```shell
make build
```
## Run
To run the tool, copy the executable to the source code directory, or to a locatoion known to PATH, then run:
```shell
./redundantptr ./...
```
