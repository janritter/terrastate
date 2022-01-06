package cmd

func buildTFInitArgs(tfInitUpgrade bool) string {
	var args string
	if tfInitUpgrade {
		args = "-upgrade"
	}
	return args
}
