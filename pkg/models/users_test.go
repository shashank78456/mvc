package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

func TestGetUserType(t *testing.T) {
	err := godotenv.Load("../../.env")

	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	got, err := GetUserType("sadmin")
	want := "superadmin"

	if err != nil {
		t.Errorf(err.Error())
	} else if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
