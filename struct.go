package main

import "fmt"

func main() {

	var tests = []struct {
		userTestRecord               bool
		scope                        string
		expectedUserTestScope        string
		expectedUserTestRecordExpiry string
		expectedOidc                 string
	}{
		{true, "test:*", "*", "2022-08-30T03:58:50.562Z", `"user":true,"scope":"*","user_expiry":"2022-08-30T03:58:50.562Z"`},
		{true, "test test:*", "*", "2022-08-30T03:58:50.562Z", `"user":true,"scope":"*","user_expiry":"2022-08-30T03:58:50.562Z"`},
		{true, "test:a test:b", "a,b", "2022-08-30T03:58:50.562Z", `"user":true,"scope":"a,b","user_expiry":"2022-08-30T03:58:50.562Z"`},
	}

	for _, tt := range tests {

		fmt.Printf(tt.expectedOidc)
		fmt.Printf("\n")

	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
