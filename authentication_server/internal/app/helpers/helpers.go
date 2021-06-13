// package helpers contain different helper functions which can be used by differen modules
package helpers

// Helper functoin to check if string is in slice. if it is returnes true. else false
func StringInSlice(teststring string, slice []string) bool {
	for _, x := range slice {
		if teststring == x {
			return true
		}
	}
	return false
}
