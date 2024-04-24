# Basic Go

## Variable and Data Type
### Declaring variable
**Define a name and type**. Here, you declare a variable with the keyword `var`, give it a name and lastly a type `string`.

```golang
var smsSendingLimit int
var costPerSMS float64
var hasPermission bool
var username string
 ```

**Define a group** of variables. It's possible to define a group of variables. Using this way of declaring means you only type the `var`keyword once.

```golang
var (
	temperature = 0.0
	isfunny = true
)
```

**Define and assign a value**. Within functions, you can use the `:=` operator, it declares and assigns at the same time.
```golang
numCars := 10
```

### Function