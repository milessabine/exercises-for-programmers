/*Enter two strings and I'll tell you if they are anagrams:
Enter the first string: note
Enter the second string: tone
"note" and "tone" are anagrams.
*/

package main

func main() {

}

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, x := range a {
		found := false

		for _, y := range b {
			if y == x {
				found = true
				break
			}
		}

		if found == false {
			return false
		}
	}

	return true
}
