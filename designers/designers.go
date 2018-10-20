package designers

import "github.com/go-modules-by-example-staging/goinfo/contributors"

func Names() []string {
	var res []string
	for _, p := range contributors.Details() {
		switch p.Name {
		case "Rob Pike", "Ken Thompson", "Robert Griesemer":
			res = append(res, p.Name)
		}
	}
	return res
}
