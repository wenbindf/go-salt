package cmd

// RunResult ...
type RunResult struct {
	Pid     int    `json:"pid"`
	Retcode int    `json:"retcode"`
	Stderr  string `json:"stderr"`
	Stdout  string `json:"stdout"`
}

// ExecCodeAll Pass in two strings, the first naming the executable language, aka -
// python2, python3, ruby, perl, lua, etc. the second string containing
// the code you wish to execute. All cmd artifacts (stdout, stderr, retcode, pid)
// will be returned.
// - cmd.exec_code_all
func (ci *Impl) ExecCodeAll(language, code string) (result map[string]*RunResult, err error) {
	return result, ci.client.RunCmd(ci.target, "cmd.exec_code_all", []interface{}{language, code}, nil, &result)
}

// ExecCode Pass in two strings, the first naming the executable language, aka -
// python2, python3, ruby, perl, lua, etc. the second string containing
// the code you wish to execute. The stdout will be returned.
// - cmd.exec_code
func (ci *Impl) ExecCode(language, code string) (result map[string]string, err error) {
	return result, ci.client.RunCmd(ci.target, "cmd.exec_code", []interface{}{language, code}, nil, &result)
}
