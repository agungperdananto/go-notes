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

## Data types

There are many data types you can use with Go. They are divided into different categories:

- **Basic types**. In this category, we find types like integers, floats (numbers with decimals) and other types like Booleans (for true/false), strings (for text) and more.
- **Composite types**. We will talk about composite types in a separate article, but they are more complex, and examples of composite types are arrays, structs and interfaces.

### Declare a variable with a type

There are two ways you can declare a variable and give it a type:

- **explicitly**, by specifying its type, for example:

   ```go
   var name string
   ```

- **implicitly**, by assigning it a value and having it been inferred:

   ```go
   numCars := 10
   ```

   In the preceding code, the data type is inferred by the value you give it. In this case, the data type becomes `string` based on the value 10.

### Function