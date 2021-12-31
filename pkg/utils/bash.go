package utils

import (
	"github.com/ContinuousEngineeringProject/ce-lib/pkg/utils/error"
	"os/exec"
)

type ArgBashCmdInput struct {
	exec string
	args []string
}

type ArgBashCmdOutput struct {
	output string
	error  error.FactoryError
}

/* ToDo: add support for run and wait execution
func BashCmdRun(input ArgBashCmdInput) (output string, err error) {

	// construct command
	cmd := EXEC.Command(input.EXEC, input.args...)

	// configure stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	// run command
	if err := cmd.Run(); err != nil {
		return "Error", err // return error
	}

	return "Success", nil // return success
}

func BashCmdWait(input ArgBashCmdInput) (output string, err error) {

	// construct command
	cmd := EXEC.Command(input.EXEC, input.args...)

	// configure stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	// start command
	if err := cmd.Start(); err != nil {
		return "Error", err // return error
	}

	// wait for command to finish
	if err := cmd.Wait(); err != nil {
		return "Error", err // return error
	}

	return "Success", nil // return success
}
*/

func BashCmdOutput(input ArgBashCmdInput) (output ArgBashCmdOutput) {

	// construct command
	cmd := exec.Command(input.exec, input.args...)

	// run command
	if cmdOutput, err := cmd.Output(); err != nil {
		output.output = "Error"
		output.error = error.NewFactoryError(error.BashPackage, err.Error())
	} else {
		output.output = string(cmdOutput)
	}

	return
}
