// auth_test.go

package auth // access the internal auth package

import (
	"errors"
	"net/http"
	"testing" // importing testing package for unit tests
)

func TestGetAPIKey(t *testing.T) { // t is the test object
	// Define all test scenarios
	testScenarios := []struct {
		scenarioName   string      // descriptive name for this test case
		inputHeaders   http.Header // what we pass to GetAPIKey()
		expectedAPIKey string      // the API key we expect back
		expectedError  error       // the error we expect (or nil)
	}{
		// Test all the different scenarios
		{
			scenarioName:   "happy path",
			inputHeaders:   http.Header{"Authorization": []string{"ApiKey abc123"}},
			expectedAPIKey: "abc123",
			expectedError:  nil,
		},
		{
			scenarioName:   "no authorization header",
			inputHeaders:   http.Header{},
			expectedAPIKey: "",
			expectedError:  ErrNoAuthHeaderIncluded,
		},
		{
			scenarioName:   "malformed header - no space after ApiKey",
			inputHeaders:   http.Header{"Authorization": []string{"ApiKeyabc123"}},
			expectedAPIKey: "",
			expectedError:  errors.New("malformed authorization header"),
		},
		{
			scenarioName:   "malformed header - less than 2 elements in slice",
			inputHeaders:   http.Header{"Authorization": []string{"ApiKey"}},
			expectedAPIKey: "",
			expectedError:  errors.New("malformed authorization header"),
		},
		{
			scenarioName:   "malformed header - not ApiKey prefix",
			inputHeaders:   http.Header{"Authorization": []string{"Bearer abc123"}},
			expectedAPIKey: "",
			expectedError:  errors.New("malformed authorization header"),
		},
	}

	// Test each scenario
	for _, currentScenario := range testScenarios {
		t.Run(currentScenario.scenarioName, func(subtestRunner *testing.T) { // run the test
			// Call the function we're testing
			actualAPIKey, actualError := GetAPIKey(currentScenario.inputHeaders)

			// Verify the API key matches what we expected
			if actualAPIKey != currentScenario.expectedAPIKey {
				subtestRunner.Errorf("expected key %q, got %q", currentScenario.expectedAPIKey, actualAPIKey)
			}

			// Verify the error matches what we expected
			if actualError != nil && currentScenario.expectedError != nil {
				if actualError.Error() != currentScenario.expectedError.Error() {
					subtestRunner.Errorf("expected error %q, got %q", currentScenario.expectedError.Error(), actualError.Error())
				}
			} else if actualError != currentScenario.expectedError {
				subtestRunner.Errorf("expected error %v, got %v", currentScenario.expectedError, actualError)
			}
		})
	}
}
