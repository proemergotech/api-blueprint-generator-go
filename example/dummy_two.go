package example

type DummyTwo struct {
	Included
	SomeParam     string      `json:"some_param"`    // example value of param
	OptionalParam *string     `json:"some_param"`    // any pointer param will be optional
	IntParam      int         `json:"int_param"`     // 10
	NestedParam   Nested      `json:"nested_param"`  //
	NestedParams  []Nested    `json:"nested_params"` //
	Anything      interface{} `json:"anything"`      // example enum (any value)
}
