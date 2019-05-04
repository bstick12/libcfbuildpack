/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package runner

import (
	"os/exec"
)

// CommandRunner is an empty struct to hang the Run method on.
type CommandRunner struct {
}

// Run makes CommandRunner satisfy the Runner interface.  This implementation delegates to exec.Command.
func (c *CommandRunner) Run(program string, args []string, dir string, env map[string]string) (string, error) {
	var cmd *exec.Cmd
	if args == nil {
		cmd = exec.Command(program)
	} else {
		cmd = exec.Command(program, args...)
	}

	if dir != "" {
		cmd.Dir = dir
	}

	if env != nil {
		for k, v := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
	}

	out, err := cmd.CombinedOutput()
	return string(out), err
}
