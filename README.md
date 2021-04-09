# go-to

[![Go Reference](https://pkg.go.dev/badge/github.com/shanbay/go-to.svg)](https://pkg.go.dev/github.com/shanbay/go-to)

`go-to` is a golang all basic types conversion/cast tool.

Example:

```go
import "github.com/shanbay/go-to"

func f() {
    sMap := map[string]string{
        // key value here...
    }

    res := ExampleStruct{
        ExampleFieldStrPtr: to.String_StringPtr(sMap["key"]),
        ExampleFieldUint32Ptr: to.Int_Uint32Ptr(100),
    }

    i64Ptr = new(int64)
    // ...

    // take value from pointer
    FuncNeedInt(to.Int64Ptr_Int(i64Ptr))
    // take value from pointer (with default value when pointer is nil)
    FuncNeedInt(to.Int64Ptr_Int(i64Ptr, to.UseDefaultEmpty))
}
```
