package cmd

// Tty Echo a string to a specific tty
// - cmd.tty
func (ci *Impl) Tty(tty, words string) (result map[string]map[string]string, err error) {
	result = map[string]map[string]string{}
	return result, ci.client.RunCmd(ci.target, "cmd.tty", []interface{}{tty, words}, nil, &result)
}

// Which Returns the path of an executable available on the minion, None otherwise
// - cmd.which
func (ci *Impl) Which(cmd string) (result map[string]string, err error) {
	return result, ci.client.RunCmd(ci.target, "cmd.which", []interface{}{cmd}, nil, &result)
}

// WhichBin Returns the first command found in a list of commands
// - cmd.which_bin
// func (ci *Impl) WhichBin(cmds string) (result map[string]string, err error) {
// 	return result, ci.client.RunCmd(ci.target, "cmd.which_bin", []interface{}{cmds}, nil, &result)
// }

// HasExec Returns true if the executable is available on the minion, false otherwise
// - cmd.has_exec
func (ci *Impl) HasExec(cmd string) (result map[string]bool, err error) {
	return result, ci.client.RunCmd(ci.target, "cmd.has_exec", []interface{}{cmd}, nil, &result)
}

// Shells Lists the valid shells on this system via the /etc/shells file
// - cmd.shells
func (ci *Impl) Shells() (result map[string][]string, err error) {
	return result, ci.client.RunCmd(ci.target, "cmd.shells", nil, nil, &result)
}
