// +build !linux

package selinux

import (
	"errors"
)

const (
	// Enforcing constant indicate SELinux is in enforcing mode
	Enforcing = 1
	// Permissive constant to indicate SELinux is in permissive mode
	Permissive = 0
	// Disabled constant to indicate SELinux is disabled
	Disabled = -1
)

var (
	// ErrMCSAlreadyExists is returned when trying to allocate a duplicate MCS.
	ErrMCSAlreadyExists = errors.New("MCS label already exists")
	// ErrEmptyPath is returned when an empty path has been specified.
	ErrEmptyPath = errors.New("empty path")
)

// Context is a representation of the SELinux label broken into 4 parts
type Context map[string]string

// SetDisabled disables selinux support for the package
func SetDisabled() {
	return
}

// SetFileLabel sets the SELinux label for this path or returns an error.
func SetFileLabel(fpath string, label string) error {
	return nil
}

// FileLabel returns the SELinux label for this path or returns an error.
func FileLabel(fpath string) (string, error) {
	return "", nil
}

/*
SetFSCreateLabel tells kernel the label to create all file system objects
created by this task. Setting label="" to return to default.
*/
func SetFSCreateLabel(label string) error {
	return nil
}

/*
FSCreateLabel returns the default label the kernel which the kernel is using
for file system objects created by this task. "" indicates default.
*/
func FSCreateLabel() (string, error) {
	return nil
}

// CurrentLabel returns the SELinux label of the current process thread, or an error.
func CurrentLabel() (string, error) {
	return "", nil
}

// PidLabel returns the SELinux label of the given pid, or an error.
func PidLabel(pid int) (string, error) {
	return "", nil
}

/*
ExecLabel returns the SELinux label that the kernel will use for any programs
that are executed by the current process thread, or an error.
*/
func ExecLabel() (string, error) {
	return "", nil
}

/*
CanonicalizeContext takes a context string and writes it to the kernel
the function then returns the context that the kernel will use.  This function
can be used to see if two contexts are equivalent
*/
func CanonicalizeContext(val string) (string, error) {
	return "", nil
}

/*
SetExecLabel sets the SELinux label that the kernel will use for any programs
that are executed by the current process thread, or an error.
*/
func SetExecLabel(label string) error {
	return nil
}

// Get returns the Context as a string
func (c Context) Get() string {
	return ""
}

// NewContext creates a new Context struct from the specified label
func NewContext(label string) Context {
	c := make(Context)
	return c
}

// ReserveLabel reserves the MLS/MCS level component of the specified label
func ReserveLabel(label string) {
	return
}

// EnforceMode returns the current SELinux mode Enforcing, Permissive, Disabled
func EnforceMode() int {
	return Disabled
}

/*
SetEnforceMode sets the current SELinux mode Enforcing, Permissive.
Disabled is not valid, since this needs to be set at boot time.
*/
func SetEnforceMode(mode int) error {
	return nil
}

/*
DefaultEnforceMode returns the systems default SELinux mode Enforcing,
Permissive or Disabled. Note this is is just the default at boot time.
EnforceMode tells you the systems current mode.
*/
func DefaultEnforceMode() int {
	return Disabled
}

/*
ReleaseLabel will unreserve the MLS/MCS Level field of the specified label.
Allowing it to be used by another process.
*/
func ReleaseLabel(label string) {
	return
}

// ROFileLabel returns the specified SELinux readonly file label
func ROFileLabel() string {
	return ""
}

/*
ContainerLabels returns an allocated processLabel and fileLabel to be used for
container labeling by the calling process.
*/
func ContainerLabels() (processLabel string, fileLabel string) {
	return "", ""
}

// SecurityCheckContext validates that the SELinux label is understood by the kernel
func SecurityCheckContext(val string) error {
	return nil
}

/*
CopyLevel returns a label with the MLS/MCS level from src label replaced on
the dest label.
*/
func CopyLevel(src, dest string) (string, error) {
	return "", nil
}

// Chcon changes the `fpath` file object to the SELinux label `label`.
// If `fpath` is a directory and `recurse`` is true, Chcon will walk the
// directory tree setting the label.
func Chcon(fpath string, label string, recurse bool) error {
	return nil
}

// DupSecOpt takes an SELinux process label and returns security options that
// can be used to set the SELinux Type and Level for future container processes.
func DupSecOpt(src string) []string {
	return nil
}

// DisableSecOpt returns a security opt that can be used to disable SELinux
// labeling support for future container processes.
func DisableSecOpt() []string {
	return []string{"disable"}
}
