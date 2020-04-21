package shopify

type ApiVersion int

const (
	VERSION_2019_07 ApiVersion = iota
	VERSION_2019_10
	VERSION_2020_01
	VERSION_2020_04
	VERSION_2020_07
	VERSION_2020_10
)

func (a ApiVersion) String() string {
	names := [...]string{
		"api/2019-07",
		"api/2019-10",
		"api/2020-01",
		"api/2020-04",
		"api/2020-07",
		"api/2020-10",
	}

	return names[a]
}
