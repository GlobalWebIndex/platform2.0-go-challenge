package main

func check(u string, p string) bool {
	pgt, found := allcreds[u]
	// fmt.Println(pgt)
	if !found {
		return false
	}
	if pgt != p {
		return false
	}
	return true
}
