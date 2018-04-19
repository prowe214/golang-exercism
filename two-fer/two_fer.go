// package twofer takes names and gives them one
package twofer

// ShareWith takes a string and gives them one
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
