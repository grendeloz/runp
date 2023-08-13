// This package provides helper code for capturing information about a
// a program's run environment.

package runp

import (
	"log"
	"os"
	"os/user"
	"strconv"
	"time"
)

var (
	Tool    string
	Version string
)

// Run environment
type RunParameters struct {
	Tool      string
	Version   string
	StartTime time.Time
	Args      []string
	UserId    int
	UserName  string
	GroupId   int
	GroupName string
	HostName  string
}

func SetVersion(v string) {
	Version = v
}

func SetTool(t string) {
	Tool = t
}

// Return a record with execution parameters
func NewRunParameters() RunParameters {
	userId := os.Getuid()
	groupId := os.Getgid()

	// Systems that use LDAP for user management bork when trying to get
	// the names to match the UID/GID numbers so we are going to silently
	// ignore errors on those functions. Caveat emptor!
	userName := ""
	tmpUserName, err := user.LookupId(strconv.Itoa(userId))
	if err == nil {
		userName = tmpUserName.Name
	}
	groupName := ""
	tmpGroupName, err := user.LookupGroupId(strconv.Itoa(groupId))
	if err == nil {
		groupName = tmpGroupName.Name
	}

	hostName, err := os.Hostname()
	if err != nil {
		// If we can't get a hostname, we will set to the error message
		hostName = err.Error()
	}

	// Setup and return RunParameters
	var run RunParameters
	run.StartTime = time.Now()
	run.Tool = Tool
	run.Version = Version
	run.Args = os.Args
	run.UserId = userId
	run.UserName = userName
	run.GroupId = groupId
	run.GroupName = groupName
	run.HostName = hostName
	return run
}

// Log logs key execution parameters at INFO loglevel.
func (rp RunParameters) Log() {
	log.Print("Tool: ", rp.Tool, ` `, rp.Version)
	log.Print("Cmdline: ", rp.Args)
	log.Print("Host: ", rp.HostName)
	log.Printf("User: %d (%s)", rp.UserId, rp.UserName)
	log.Printf("Group: %d (%s)", rp.GroupId, rp.GroupName)
}

func (rp RunParameters) LogFinish() {
	end := time.Now()
	elapsed := end.Sub(rp.StartTime)
	log.Print("Elapsed time: ", elapsed)
}
