package types

type PushEventGithub struct {
	PushID        int
	Size          int
	Distinct_Size int
	Ref           string
	Head          string
	Before        string
	Commits       []struct {
		ID      string
		Message string
		Author  struct {
			Name  string
			Email string
		}
		URL      string
		Distinct bool
	}
}
