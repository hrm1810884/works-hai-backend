package history

type History struct {
	HistoryId string
	Version   Version
}

func NewHistory(version int) *History {
	createdVersion := NewVersion(version)
	return &History{
		HistoryId: createdVersion.CreateId(),
		Version:   *createdVersion,
	}
}

func (h *History) GetHistoryId() string {
	return h.HistoryId
}

func (h *History) GetVersion() int {
	return h.Version.GetVersion()
}
