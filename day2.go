package main

func BoolToWord(word bool) string {
	if word == true {
		return "Yes"
	} else if word == false {
		return "No"
	}

	return "Insert either true or false"
}
