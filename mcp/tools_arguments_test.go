package mcp

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallToolRequestWithMapArguments(t *testing.T) {
	// Create a request with map arguments
	req := CallToolRequest{}
	req.Params.Name = "test-tool"
	req.Params.Arguments = map[string]any{
		"key1": "value1",
		"key2": 123,
	}

	// Test GetArguments
	args := req.GetArguments()
	assert.Equal(t, "value1", args["key1"])
	assert.Equal(t, 123, args["key2"])

	// Test GetRawArguments
	rawArgs := req.GetRawArguments()
	mapArgs, ok := rawArgs.(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, "value1", mapArgs["key1"])
	assert.Equal(t, 123, mapArgs["key2"])
}

func TestCallToolRequestWithNonMapArguments(t *testing.T) {
	// Create a request with non-map arguments
	req := CallToolRequest{}
	req.Params.Name = "test-tool"
	req.Params.Arguments = "string-argument"

	// Test GetArguments (should return empty map)
	args := req.GetArguments()
	assert.Empty(t, args)

	// Test GetRawArguments
	rawArgs := req.GetRawArguments()
	strArg, ok := rawArgs.(string)
	assert.True(t, ok)
	assert.Equal(t, "string-argument", strArg)
}

func TestCallToolRequestWithStructArguments(t *testing.T) {
	// Create a custom struct
	type CustomArgs struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	// Create a request with struct arguments
	req := CallToolRequest{}
	req.Params.Name = "test-tool"
	req.Params.Arguments = CustomArgs{
		Field1: "test",
		Field2: 42,
	}

	// Test GetArguments (should return empty map)
	args := req.GetArguments()
	assert.Empty(t, args)

	// Test GetRawArguments
	rawArgs := req.GetRawArguments()
	structArg, ok := rawArgs.(CustomArgs)
	assert.True(t, ok)
	assert.Equal(t, "test", structArg.Field1)
	assert.Equal(t, 42, structArg.Field2)
}

func TestCallToolRequestJSONMarshalUnmarshal(t *testing.T) {
	// Create a request with map arguments
	req := CallToolRequest{}
	req.Params.Name = "test-tool"
	req.Params.Arguments = map[string]any{
		"key1": "value1",
		"key2": 123,
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	assert.NoError(t, err)

	// Unmarshal from JSON
	var unmarshaledReq CallToolRequest
	err = json.Unmarshal(data, &unmarshaledReq)
	assert.NoError(t, err)

	// Check if arguments are correctly unmarshaled
	args := unmarshaledReq.GetArguments()
	assert.Equal(t, "value1", args["key1"])
	assert.Equal(t, float64(123), args["key2"]) // JSON numbers are unmarshaled as float64
}