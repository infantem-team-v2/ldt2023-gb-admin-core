package test

import (
	"bytes"
	"encoding/json"
	_ "gb-auth-gate/internal/auth/interface"
	"gb-auth-gate/internal/auth/model"
	_ "gb-auth-gate/internal/auth/usecase"
	"gb-auth-gate/internal/pkg/common"
	"gb-auth-gate/internal/pkg/server"
	"gb-auth-gate/pkg/thttp"
	"github.com/stretchr/testify/assert"
	netHttp "net/http"
	"testing"
)

// ======================SIGN_UP========================//
func TestSignUp(t *testing.T) {
	tests := map[string]model.SignUpRequest{
		"test1": {},
		"test2": {},
		"test3": {}}

	expected := map[string]model.SignUpResponse{
		"test1": {
			Response: common.Response{
				Message:      "Created",
				InternalCode: 201,
			},
		},
		"test2": {
			Response: common.Response{
				Message:      "User already exists!",
				InternalCode: 409,
			},
		},
		"test3": {
			Response: common.Response{
				Message:      "Unauthorized Error",
				InternalCode: 401,
			},
		},
	}
	serverTest := server.NewServer().MapHandlers()

	for i, test := range tests {
		body, err := json.Marshal(test)
		req, _ := netHttp.NewRequest(thttp.POST, "/auth/sign/up", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp, err := serverTest.App.Test(req, -1)
		assert.Equal(t, expected[i].InternalCode, resp.StatusCode)
	}
}

// =========================SIGN_IN==========================//
func TestSignIn(t *testing.T) {
	tests := map[string]model.SignInRequest{
		"test1": {
			Email:    "test@test.com",
			Password: "Password1@",
		},
		"test2": {
			Email:    "aga@da.verno",
			Password: "shutka",
		},
		"test3": {
			Email:    "net.chto.eto",
			Password: "",
		}}

	expected := map[string]model.SignInResponse{
		"test1": {
			Response: common.Response{
				Message:      "Success",
				InternalCode: 200,
			},
		},
		"test2": {
			Response: common.Response{
				Message:      "Not found!",
				InternalCode: 404,
			},
		},
		"test3": {
			Response: common.Response{
				Message:      "Unauthorized Error",
				InternalCode: 401,
			},
		},
	}
	serverTest := server.NewServer().MapHandlers()

	for i, test := range tests {
		body, err := json.Marshal(test)
		req, _ := netHttp.NewRequest(thttp.POST, "/auth/sign/in", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp, err := serverTest.App.Test(req, -1)
		assert.Equal(t, expected[i].InternalCode, resp.StatusCode)
	}
}

// ==========================VALIDATE_EMAIL=========================//
func TestValidateEmail(t *testing.T) {
	tests := map[string]model.EmailValidateRequest{
		"test1": {
			Code: 123456,
		},
		"test2": {
			Code: 123,
		},
		"test3": {
			Code: 232323232,
		}}

	expected := map[string]model.EmailValidateResponse{
		"test1": {
			Response: common.Response{
				Message:      "Accepted",
				InternalCode: 202,
			},
		},
		"test2": {
			Response: common.Response{
				Message:      "Request validation failed",
				InternalCode: 400,
			},
		},
		"test3": {
			Response: common.Response{
				Message:      "Request validation failed",
				InternalCode: 400,
			},
		},
	}
	serverTest := server.NewServer().MapHandlers()

	for i, test := range tests {
		body, err := json.Marshal(test)
		req, _ := netHttp.NewRequest(thttp.POST, "/auth/email/validate", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp, err := serverTest.App.Test(req, -1)
		assert.Equal(t, expected[i].InternalCode, resp.StatusCode)
	}
}
