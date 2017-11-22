api-blueprint-generator-go
========

## Overview

This tool can be used to generate a section of your API.md file. It only generates mson data structures.

## Development state

This project is in development stage. It is mainly used for private purposes, meaning features will not be added unless I need them or is requested. Pull requests are welcome. 

## Installation

`go get github.com/proemergotech/api-blueprint-generator-go`

## Usage

`api-blueprint-generator-go [source .go files glob] [destionation API.md file]`

Note that you can safely extend and regenerate your API.md file, as only content between the placeholders
``` ###### `# Generated docs start` ``` and ``` ###### `# Generated docs end` ``` will be replaced.

## Example

`api-blueprint-generator-go example/dummy*.go example/API.md`

### Source files

#### dummy_one.go
```go
package example

type Dummy struct {
	Included
	SomeParam     string      `json:"some_param"`    // example value of param
	OptionalParam *string     `json:"some_param"`    // any pointer param will be optional
	IntParam      int         `json:"int_param"`     // 10
	BoolParam     bool        `json:"bool_param"`    // true
	NestedParam   Nested      `json:"nested_param"`  //
	NestedParams  []Nested    `json:"nested_params"` //
	Anything      interface{} `json:"anything"`      // example enum (any value)
}

type Included struct {
	IncludedParam string `json:"included_param"` // example value of included param
}

type Nested struct {
	OtherParam string `json:"other_param"` // example value of other param
}
```

#### dummy_two.go
```go
package example

type DummyTwo struct {
	Included
	SomeParam     string      `json:"some_param"`    // example value of param
	OptionalParam *string     `json:"some_param"`    // any pointer param will be optional
	IntParam      int         `json:"int_param"`     // 10
	BoolParam     bool        `json:"bool_param"`    // true
	NestedParam   Nested      `json:"nested_param"`  //
	NestedParams  []Nested    `json:"nested_params"` //
	Anything      interface{} `json:"anything"`      // example enum (any value)
}
```

### Generated API.md

```markdown
# Example generated API.md file

You can write any content here.

###### `# Generated docs start`

### Dummy
+ Include Included
+ some_param: `example value of param` (string, required)
+ some_param: `any pointer param will be optional` (string, optional)
+ int_param: `10` (number, required)
+ nested_param (Nested, required)
+ nested_params (array[Nested], required)
+ anything: `example enum (any value)` (enum, required)

### Included
+ included_param: `example value of included param` (string, required)

### Nested
+ other_param: `example value of other param` (string, required)

### DummyTwo
+ Include Included
+ some_param: `example value of param` (string, required)
+ some_param: `any pointer param will be optional` (string, optional)
+ int_param: `10` (number, required)
+ nested_param (Nested, required)
+ nested_params (array[Nested], required)
+ anything: `example enum (any value)` (enum, required)

###### `# Generated docs end`
```
