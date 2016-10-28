package cmd

// Script Download a script from a remote location and execute the script locally.
// The script can be located on the salt master file server or on an HTTP/FTP
// server.
// The script will be executed directly, so it can be written in any available
// programming language.
// - cmd.script
func (ci *Impl) Script(source, args string) (result map[string]*RunResult, err error) {
	return result, ci.client.RunCmd(ci.target, "cmd.script", []interface{}{source, args}, ci.kwarg, &result)
}

// ScriptRetcode Only evaluate the script return code and do not block for terminal output
// - cmd.script_retcode
func (ci *Impl) ScriptRetcode(source, args string) (result map[string]int, err error) {
	return result, ci.client.RunCmd(ci.target, "cmd.script_retcode", []interface{}{source, args}, ci.kwarg, &result)
}
