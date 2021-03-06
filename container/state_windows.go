package container

// setFromExitStatus is a platform specific helper function to set the state
// based on the ExitStatus structure.
func (s *State) setFromExitStatus(exitStatus *ExitStatus) {
	s.ExitCode = exitStatus.ExitCode
}
