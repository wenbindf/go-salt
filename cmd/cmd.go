package cmd

import salt "github.com/xuguruogu/go-salt"

// Cmd ...
type Cmd interface {
	//cmd
	Run(cmd string, param ...*Param) (result map[string]string, err error)
	//exec
	ExecCode(language, code string) (result map[string]string, err error)
	ExecCodeAll(language, code string) (result map[string]*ExeResult, err error)
	//utils
	Tty(tty, words string) (result map[string]map[string]string, err error)
	Which(cmd string) (result map[string]string, err error)
	WhichBin(cmds string) (result map[string]string, err error)
	HasExec(cmd string) (result map[string]bool, err error)
	Shells() (result map[string][]string, err error)
}

// Impl ...
type Impl struct {
	client salt.Client
	target string
}

// New ...
func New(target string, client salt.Client) Cmd {
	return &Impl{target: target, client: client}
}

// Param ...
type Param struct {
	Env               map[string]string `json:"env,omitempty"`                 //environment variables
	Cwd               string            `json:"cwd,omitempty"`                 //working directory
	Stdin             string            `json:"stdin,omitempty"`               //stdin
	Runas             string            `json:"runas,omitempty"`               //User to run script as
	Password          string            `json:"password,omitempty"`            //Windows only
	Shell             string            `json:"shell,omitempty"`               //Shell to execute under
	PythonShell       bool              `json:"python_shell,omitempty"`        //Set to True to use shell features, such as pipes or redirection
	CleanEnv          bool              `json:"clean_env,omitempty"`           //Attempt to clean out all other shell environment variables and set only those provided in the 'env' argument to this function.
	Template          string            `json:"template,omitempty"`            //Currently jinja, mako, and wempy are supported
	Rstrip            bool              `json:"rstrip,omitempty"`              //Strip all whitespace off the end of output before it is returned
	Umask             string            `json:"umask,omitempty"`               //The umask (in octal) to use when running the command
	OutputLogLevel    string            `json:"output_loglevel,omitempty"`     //Control the loglevel at which the output from the command is logged
	Timeout           int               `json:"timeout,omitempty"`             //A timeout in seconds for the executed process to return.
	UseVt             bool              `json:"use_vt,omitempty"`              //Use VT utils (saltstack) to stream the command output more interactively to the console and the logs. This is experimental.
	ResetSystemLocale bool              `json:"reset_system_locale,omitempty"` //Resets the system locale
	IgnoreRetcode     bool              `json:"ignore_retcode,omitempty"`      //Ignore the return code
	Saltenv           string            `json:"saltenv,omitempty"`             //The salt environment to use. Default is 'base'
}

// ExeParam ...
type ExeParam struct {
	*Param
	Cmd string `json:"cmd,omitempty"` //param str cmd: The command to run
}

// Run Execute the passed command and return the output as a string
// at least 1 argument
// - cmd.run
func (ci *Impl) Run(cmd string, param ...*Param) (result map[string]string, err error) {
	kwarg := ExeParam{
		Cmd: cmd,
		Param: func() *Param {
			if len(param) != 0 {
				return param[0]
			}
			return nil
		}(),
	}

	return result, ci.client.RunCmd(ci.target, "cmd.run", nil, kwarg, &result)
}

// Retcode Execute a shell command and return the command's return code.
// at least 1 argument
// - cmd.retcode
func (ci *Impl) Retcode(cmd string, param ...*Param) {

}

// RunStderr Execute a command and only return the standard error
// at least 1 argument
// - cmd.run_stderr
func (i *Impl) RunStderr(cmd string, param ...*Param) {

}

// RunStdout Execute a command and only return the standard out
// at least 1 argument
// - cmd.run_stdout
func (i *Impl) RunStdout(cmd string, param ...*Param) {

}

// RunAll Execute the passed command and return a dict of return data
// at least 1 argument
// - cmd.run_all
func (i *Impl) RunAll(cmd string, param ...*Param) {

}

// RunBg Execute the passed command in the background and return it's PID
// at least 1 argument
// - cmd.run_bg
func (i *Impl) RunBg(cmd string, param ...*Param) {

}

// ChrootExeParam ...
type ChrootExeParam struct {
	Param
	Root string `json:"root,omitempty"` //Path to the root of the jail to use, see chroot
}

// RunChroot wrapped within a chroot, with dev and proc mounted in the chroot
// at least 2 argument
// - cmd.run_chroot
func (i *Impl) RunChroot(param *ChrootExeParam) {

}

// Powershell Execute the passed PowerShell command and return the output as a string.
// at least 1 argument
// - cmd.powershell
func (i *Impl) Powershell(cmd string, param ...*Param) {

}

// Shell Execute the passed command and return the output as a string
// at least 1 argument
// - cmd.shell
func (i *Impl) Shell(cmd string, param ...*Param) {

}
