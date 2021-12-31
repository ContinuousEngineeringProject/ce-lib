package utils

import (
	"testing"
)

type wantType struct {
	output      string
	errorExists bool
}
type testData struct {
	name string
	args ArgBashCmdInput
	want wantType
}

// Define test data
var testCases = []testData{
	{
		name: "Execute a command to return an output",
		args: ArgBashCmdInput{
			exec: "echo",
			args: []string{"Hello World"},
		},
		want: wantType{
			output:      "Hello World\n",
			errorExists: false,
		},
	},
	{
		name: "Execute a command to return an error",
		args: ArgBashCmdInput{
			exec: "ech",
			args: []string{"Hello World\n"},
		},
		want: wantType{
			output:      "Error",
			errorExists: true,
		},
	},
	{
		name: "Execute a script to return an output",
		args: ArgBashCmdInput{
			exec: "/bin/sh",
			args: []string{"./testdata/BashCmdToReturnAnOutput.sh"},
		},
		want: wantType{
			output:      "Hello World\n",
			errorExists: false,
		},
	},
	{
		name: "Execute a script to return an error",
		args: ArgBashCmdInput{
			exec: "/bin/sh",
			args: []string{"./testdata/BashCmdToReturnAnError.sh"},
		},
		want: wantType{
			output:      "Error",
			errorExists: true,
		},
	},
}

// Test BashCmdOutput
//
func TestBashCmdOutput(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			testStatus := BashCmdOutput(tt.args)
			if testStatus.error.Exists() != tt.want.errorExists || testStatus.output != tt.want.output {
				t.Errorf("%v test failed, expected output: %v and error: %v, got output: %v and error: %v \n %v", tt.name, tt.want.output, tt.want.errorExists, testStatus.output, testStatus.error.Exists(), testStatus.error.ErrorString())
			}
		})
	}
}
