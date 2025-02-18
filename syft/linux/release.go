package linux

// Release represents Linux Distribution release information as specified from https://www.freedesktop.org/software/systemd/man/os-release.html
type Release struct {
	PrettyName       string   // A pretty operating system name in a format suitable for presentation to the user.
	Name             string   // identifies the operating system, without a version component, and suitable for presentation to the user.
	ID               string   // identifies the operating system, excluding any version information and suitable for processing by scripts or usage in generated filenames.
	IDLike           []string // list of operating system identifiers in the same syntax as the ID= setting. It should list identifiers of operating systems that are closely related to the local operating system in regards to packaging and programming interfaces.
	Version          string   // identifies the operating system version, excluding any OS name information, possibly including a release code name, and suitable for presentation to the user.
	VersionID        string   // identifies the operating system version, excluding any OS name information or release code name, and suitable for processing by scripts or usage in generated filenames.
	Variant          string   // identifies a specific variant or edition of the operating system suitable for presentation to the user.
	VariantID        string   // identifies a specific variant or edition of the operating system. This may be interpreted by other packages in order to determine a divergent default configuration.
	HomeURL          string
	SupportURL       string
	BugReportURL     string
	PrivacyPolicyURL string
	CPEName          string // A CPE name for the operating system, in URI binding syntax
}

func (r *Release) String() string {
	if r == nil {
		return "unknown"
	}
	if r.PrettyName != "" {
		return r.PrettyName
	}
	if r.Name != "" {
		return r.Name
	}
	if r.Version != "" {
		return r.ID + " " + r.Version
	}

	return r.ID + " " + r.VersionID
}
