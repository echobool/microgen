# Microgen

Tool to generate microservices, based on [go-kit](https://gokit.io/), by specified service interface.

## Usage
``` sh
microgen [OPTIONS]
```
### Options

| Name        | Default          | Description                                                                   |
|:------------|:-----------------|:------------------------------------------------------------------------------|
| -file*      |                  | Relative path to source file with service interface                           |
| -interface* |                  | Service interface name in source file                                         |
| -out        | writes to stdout | Relative or absolute path to directory, where you want to see generated files |
| -package*   |                  | Package path of your service interface source file                            |
| -debug      | false            | Display some debug information                                                |

\* __Required option__

### Interface declaration rules
For correct generation, please, follow rules below.

* All interface arguments and results should be named.
* First argument should be of type `context.Context` (from [standard library](https://golang.org/pkg/context/)).
* Result arguments should contain at least one variable of `error` type.
* Argument's and result's names should be different (name duplicating unacceptable).
* [Some names](#not-allowed-names) are not allowed to be an argument or result.

#### Not allowed names:
```
req
resp
```