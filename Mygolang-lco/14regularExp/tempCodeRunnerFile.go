sempleRegx := regexp.MustCompile("abd|efg")

	match2 := sempleRegx.Match([]byte("abd"))
	fmt.Println(match2)
	match2 = sempleRegx.Match([]byte("efg"))
	fmt.Println(match2)
	match2 = sempleRegx.Match([]byte("abdefg"))
	fmt.Println(match2)
	match2 = sempleRegx.Match([]byte("abc"))
	fmt.Println(match2)