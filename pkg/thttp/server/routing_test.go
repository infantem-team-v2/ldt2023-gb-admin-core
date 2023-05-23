package server

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestParseRoutes Main Testing file
func TestParseRoutes(t *testing.T) {
	// Testing opening and reading file
	t.Run("Testing full right file", TestFullFile)
	t.Run("Testing empty file", TestEmptyFile)
	t.Run("Testing no file", TestNoFile)

	t.Run("Testing map routes", TestMapRoutes)
}

func TestFullFile(t *testing.T) {
	testCase := map[string]RouteMapping{
		"SignIn": {
			HttpMethod: "POST",
			Route:      "/sign/in",
		},
		"SignUp": {
			HttpMethod: "POST",
			Route:      "/sign/up",
		},
		"VendorAuth": {
			HttpMethod: "POST",
			Route:      "/",
		},
	}
	res, err := ParseRoutes("/test/full.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testCase, res, "Test result")

	t.Logf("Result: %v", res)
}

func TestEmptyFile(t *testing.T) {
	res, err := ParseRoutes("/test/empty.yaml")
	if res != nil {
		t.Fatal("not empty map returned from empty file")
	}

	assert.Equal(t, err.Error(), errors.Wrap(ErrEmptyFile, fmt.Sprintf(wrapLabel, "/test/empty.yaml")).Error(), "Error assert result")

}

func TestNoFile(t *testing.T) {
	res, err := ParseRoutes("/test/no-file.yaml")
	if res != nil {
		t.Fatal("not empty map returned from no file")
	}
	assert.NotEqualf(t, err, nil, "Error from no file")

}

func TestMapRoutes(t *testing.T) {

}
