package conv

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSnakeToCamel(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "hello",
		},
		{
			Input:  "hello_world",
			Output: "helloWorld",
		},
		{
			Input:  "hello_world_",
			Output: "helloWorld",
		},
		{
			Input:  "hello_world_id",
			Output: "helloWorldID",
		},
	}

	for _, tc := range tests {
		if diff := cmp.Diff(tc.Output, ToCamel(tc.Input)); diff != "" {
			t.Fatal(diff)
		}
	}
}

func TestSnakeToClass(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "Hello",
		},
		{
			Input:  "hello_world",
			Output: "HelloWorld",
		},
		{
			Input:  "hello_world_",
			Output: "HelloWorld",
		},
		{
			Input:  "hello_world_id",
			Output: "HelloWorldID",
		},
	}

	for _, tc := range tests {
		if diff := cmp.Diff(tc.Output, ToCamel(tc.Input)); diff == "" {
			t.FailNow()
		}
	}
}

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "hello",
			Output: "hello",
		},
		{
			Input:  "helloWorld",
			Output: "hello_world",
		},
		{
			Input:  "helloWorldID",
			Output: "hello_world_id",
		},
	}

	for _, tc := range tests {
		if diff := cmp.Diff(tc.Output, ToSnake(tc.Input)); diff != "" {
			t.Fatal(diff)
		}
	}
}
