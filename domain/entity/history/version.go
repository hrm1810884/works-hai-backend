package history

import "strconv"

type Version struct {
	Version int
}

func NewVersion(version int) *Version {
	return &Version{
		Version: version,
	}
}

func (v *Version) GetVersion() int {
	return v.Version
}

func (v *Version) CreateId() string {
	return strconv.Itoa(v.GetVersion())
}

func (v *Version) GetNextVersion() *Version {
	return NewVersion(v.Version + 1)
}

func (v *Version) GetPreviousVersion() *Version {
	if v.Version <= 0 {
		return nil
	}
	return NewVersion(v.Version - 1)
}
