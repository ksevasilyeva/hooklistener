package provider

type Provider int

const (
	Unknown Provider = iota
	GitHub
	Bitbucket
)

func (p Provider) String() string {
	switch p {
	case GitHub:
		return "GitHub"
	case Bitbucket:
		return "Bitbucket"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}
