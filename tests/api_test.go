package tests

import (
	"testing"
)

func apiTest(t *testing.T) {
	t.Log("API test executed")
}

func TestAPIRoutes(t *testing.T) {
	t.Run("TestAPI", apiTest)
}
