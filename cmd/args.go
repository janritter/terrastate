package cmd

func buildTFInitArgs(tfInitUpgrade bool) []string {
	if tfInitUpgrade {
		return []string{"-upgrade"}
	}
	return nil
}
