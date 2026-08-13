package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestGetApiKeyReturnKey(t *testing.T) {
	type testCase struct {
		headers        http.Header
		expectedApiKey string
		err            error
	}

	runCases := []testCase{
		{headers: http.Header{"Authorization": []string{"ApiKey 123"}}, expectedApiKey: "123", err: nil},
		{headers: http.Header{}, expectedApiKey: "", err: errors.New("no authorization header included")},
		{headers: http.Header{"Authorization": []string{"ApiKey"}}, expectedApiKey: "", err: errors.New("malformed authorization header")},
	}

	passedCount := 0
	failedCount := 0

	for _, tc := range runCases {
		apikey, err := GetAPIKey(tc.headers)
		errStr := ""

		if err != nil {
			errStr = err.Error()
		}

		expectedErrStr := ""

		if tc.err != nil {
			expectedErrStr = tc.err.Error()
		}

		if apikey != tc.expectedApiKey || errStr != expectedErrStr {
			failedCount++
			failTc := fmt.Sprintf(`---------------------------------
Inputs: (%v)
Expecting: (%v, %v)
Actual: (%v, %v)
Fail
			`, tc.headers, tc.expectedApiKey, tc.err, apikey, err)
			t.Fatal(failTc)
		} else {
			passedCount++
			fmt.Printf(`---------------------------------
Inputs: (%v)
Expecting: (%v, %v)
Actual: (%v, %v)
Pass
			`, tc.headers, tc.expectedApiKey, tc.err, apikey, err)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passedCount, failedCount)
}
