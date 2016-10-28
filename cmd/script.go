package cmd

// ScriptExeParam ...
type ScriptExeParam struct {
	*Param
	Source string `json:"source,omitempty"` // The location of the script to download. If the file is located on the master in the directory named spam, and is called eggs, the source string is salt://spam/eggs
	Args   string `json:"args,omitempty"`   //String of command line args to pass to the script
}

// Script Download a script from a remote location and execute the script locally.
// The script can be located on the salt master file server or on an HTTP/FTP
// server.
// The script will be executed directly, so it can be written in any available
// programming language.
// - cmd.script
func (ci *Impl) Script(source, args string, param ...*Param) (result map[string]*ExeResult, err error) {
	kwarg := ScriptExeParam{
		Source: source,
		Args:   args,
		Param: func() *Param {
			if len(param) != 0 {
				return param[0]
			}
			return nil
		}(),
	}
	return result, ci.client.RunCmd(ci.target, "cmd.script", nil, kwarg, &result)
}

// ScriptRetcode Only evaluate the script return code and do not block for terminal output
// - cmd.script_retcode
func (ci *Impl) ScriptRetcode(source, args string, param ...*Param) (result map[string]int, err error) {
	kwarg := ScriptExeParam{
		Source: source,
		Args:   args,
		Param: func() *Param {
			if len(param) != 0 {
				return param[0]
			}
			return nil
		}(),
	}
	return result, ci.client.RunCmd(ci.target, "cmd.script_retcode", nil, kwarg, &result)
}
