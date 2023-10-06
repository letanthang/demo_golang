package constant

type PageCagegory int

const (
	HealthCategory PageCagegory = iota + 1
	SportCategory
	EconomyCategory
	ScienceCategory
)

func (c PageCagegory) String() string{
	switch c {
	case HealthCategory:
		return "Health"
	case SportCategory:
		return "Sport"
	case EconomyCategory:
		return "Economy"
	case ScienceCategory:
		return "Science"
	default:
		return ""
	}
}