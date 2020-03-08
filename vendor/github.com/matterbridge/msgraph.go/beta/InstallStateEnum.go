// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// InstallState undocumented
type InstallState int

const (
	// InstallStateVNotApplicable undocumented
	InstallStateVNotApplicable InstallState = 0
	// InstallStateVInstalled undocumented
	InstallStateVInstalled InstallState = 1
	// InstallStateVFailed undocumented
	InstallStateVFailed InstallState = 2
	// InstallStateVNotInstalled undocumented
	InstallStateVNotInstalled InstallState = 3
	// InstallStateVUninstallFailed undocumented
	InstallStateVUninstallFailed InstallState = 4
	// InstallStateVUnknown undocumented
	InstallStateVUnknown InstallState = 5
)

// InstallStatePNotApplicable returns a pointer to InstallStateVNotApplicable
func InstallStatePNotApplicable() *InstallState {
	v := InstallStateVNotApplicable
	return &v
}

// InstallStatePInstalled returns a pointer to InstallStateVInstalled
func InstallStatePInstalled() *InstallState {
	v := InstallStateVInstalled
	return &v
}

// InstallStatePFailed returns a pointer to InstallStateVFailed
func InstallStatePFailed() *InstallState {
	v := InstallStateVFailed
	return &v
}

// InstallStatePNotInstalled returns a pointer to InstallStateVNotInstalled
func InstallStatePNotInstalled() *InstallState {
	v := InstallStateVNotInstalled
	return &v
}

// InstallStatePUninstallFailed returns a pointer to InstallStateVUninstallFailed
func InstallStatePUninstallFailed() *InstallState {
	v := InstallStateVUninstallFailed
	return &v
}

// InstallStatePUnknown returns a pointer to InstallStateVUnknown
func InstallStatePUnknown() *InstallState {
	v := InstallStateVUnknown
	return &v
}