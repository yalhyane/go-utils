# Go-utils

**Go-utils** is a collection of helper functions for common tasks related to:

* **Slices:** Handle and manipulate slices more easily with functions like `DeleteValue`, `FilterFunc`, `IndexOf`...
* **Maps:** Work with maps efficiently using utilities like `MergeMaps`, `GetValues`, `GetKeys`...
* **Math:** Simplify mathematical operations with functions like Generic `Min` and `Max`...


**Usage Examples:**

```go
import (
  "github.com/yalhyane/go-utils/map_utils"
  "github.com/yalhyane/go-utils/slice_utils"
  "github.com/yalhyane/go-utils/math_utils"
)
// Fetch all values from a map
m := make(map[string]int)
m["a"] = 1
m["b"] = 2
values := map_utils.GetValues(m) // values will be []int{1, 2}
fmt.Println(values)

// Filter a slice
s := []string{"a", "b", "d", "c"}
s2 := slice_utils.FilterFunc(s, func(s string) bool {
    return s != "a"
}) // s2 will be []string{"b", "d", "c"}
fmt.Println(s2)

// Generic Min and Max
i2 := 2
i3 := 3
minNumber := math_utils.Min(i2, i3)
fmt.Println(minNumber)
```
