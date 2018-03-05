package main

func Check(link string) string {
	link = ValidateURL(link)

	if IsURL(link) {
		if IsIP(link) {
			if IsPrivate(link) {
				return GenerateErrorJSON(1)
			} else {
				return Fetch(link)
			}
		} else {
			return Fetch(link)
		}
	} else {
		return GenerateErrorJSON(0)
	}
}
