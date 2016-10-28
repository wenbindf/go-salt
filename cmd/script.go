package cmd

// ScriptExeParam ...
type ScriptExeParam struct {
	Param
	Source string `json:"source,omitempty"` // The location of the script to download. If the file is located on the master in the directory named spam, and is called eggs, the source string is salt://spam/eggs
	Args   string `json:"args,omitempty"`   //String of command line args to pass to the script
}

// Script Download a script from a remote location and execute the script locally.
// The script can be located on the salt master file server or on an HTTP/FTP
// server.
// The script will be executed directly, so it can be written in any available
// programming language.
// - cmd.script
func (i *Impl) Script(param ScriptExeParam) (map[string]string, error) {
	return nil, nil
}

// ScriptRetcode Only evaluate the script return code and do not block for terminal output
// - cmd.script_retcode
func (i *Impl) ScriptRetcode(param ScriptExeParam) {

}
