package example

type Dummy struct {
	Included
	SomeParam       string      `json:"some_param"`    // example value of param
	OptionalParam   *string     `json:"some_param"`    // any pointer param will be optional
	IntParam        int         `json:"int_param"`     // 10
	BoolParam       bool        `json:"bool_param"`    // true
	NestedParam     Nested      `json:"nested_param"`  //
	NestedParams    []Nested    `json:"nested_params"` //
	Anything        interface{} `json:"anything"`      // example enum (any value)
	InterfaceAsLast interface{}
}

type Included struct {
	IncludedParam string `json:"included_param"` // example value of included param
}

type Nested struct {
	OtherParam string `json:"other_param"` // example value of other param
}
