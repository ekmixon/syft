package pkg

import (
	"sort"

	"github.com/anchore/syft/syft/file"

	"github.com/anchore/packageurl-go"
	"github.com/anchore/syft/syft/linux"
	"github.com/scylladb/go-set/strset"
)

const DpkgDBGlob = "**/var/lib/dpkg/{status,status.d/**}"

var _ FileOwner = (*DpkgMetadata)(nil)

// DpkgMetadata represents all captured data for a Debian package DB entry; available fields are described
// at http://manpages.ubuntu.com/manpages/xenial/man1/dpkg-query.1.html in the --showformat section.
type DpkgMetadata struct {
	Package       string           `mapstructure:"Package" json:"package"`
	Source        string           `mapstructure:"Source" json:"source"`
	Version       string           `mapstructure:"Version" json:"version"`
	SourceVersion string           `mapstructure:"SourceVersion" json:"sourceVersion"`
	Architecture  string           `mapstructure:"Architecture" json:"architecture"`
	Maintainer    string           `mapstructure:"Maintainer" json:"maintainer"`
	InstalledSize int              `mapstructure:"InstalledSize" json:"installedSize"`
	Files         []DpkgFileRecord `json:"files"`
}

// DpkgFileRecord represents a single file attributed to a debian package.
type DpkgFileRecord struct {
	Path         string       `json:"path"`
	Digest       *file.Digest `json:"digest,omitempty"`
	IsConfigFile bool         `json:"isConfigFile"`
}

// PackageURL returns the PURL for the specific Debian package (see https://github.com/package-url/purl-spec)
func (m DpkgMetadata) PackageURL(distro *linux.Release) string {
	if distro == nil {
		return ""
	}
	pURL := packageurl.NewPackageURL(
		// TODO: replace with `packageurl.TypeDebian` upon merge of https://github.com/package-url/packageurl-go/pull/21
		// TODO: or, since we're now using an Anchore fork of this module, we could do this sooner.
		"deb",
		distro.ID,
		m.Package,
		m.Version,
		packageurl.Qualifiers{
			{
				Key:   "arch",
				Value: m.Architecture,
			},
		},
		"")
	return pURL.ToString()
}

func (m DpkgMetadata) OwnedFiles() (result []string) {
	s := strset.New()
	for _, f := range m.Files {
		if f.Path != "" {
			s.Add(f.Path)
		}
	}
	result = s.List()
	sort.Strings(result)
	return
}
