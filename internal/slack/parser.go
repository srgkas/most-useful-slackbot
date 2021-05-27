package slack

type Hotfix struct {
	Project string
	Tag string
}

type HotfixMessage struct {
	// this might be extended with extra fields
	fixes []*Hotfix
}

func (m *HotfixMessage) GetFixes() []*Hotfix {
	return m.fixes
}

// ParseHotfixMessage parses message from hotfix
func ParseHotfixMessage(message string) (*HotfixMessage, error) {
	// Might be several projects with tags
	// stub
	// parse logic

	var fixes []*Hotfix
	fix := newHotFix("srgkas/most-useful-slackbot", "v0.0.1")
	fixes = append(fixes, fix)

	return newHotfixMessage(fixes), nil
}

func newHotfixMessage(fixes []*Hotfix) *HotfixMessage {
	return &HotfixMessage{
		fixes: fixes,
	}
}

func newHotFix(project string, tag string) *Hotfix {
	return &Hotfix{
		Project: project,
		Tag: tag,
	}
}
