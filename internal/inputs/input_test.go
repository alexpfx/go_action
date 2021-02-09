package inputs

/*func TestRofiInputReader_Read(t *testing.T) {
	tests := []struct {
		inputList   *ResolverConfig
		readValues  string
		expected    []string
		errExpected string
	}{
		{
			inputList: &ResolverConfig{
				ArgSep: " ",
				Keys:   nil,
			},
			readValues:  "a b c d",
			expected:    []string{"a", "b", "c", "d"},
			errExpected: "",
		},
		{
			inputList: &ResolverConfig{
			Keys:   []string{"-t", "-x"},
			ArgSep: "|",
		}, readValues: "a|b|123", expected: []string{"-t", "a", "-x", "b", "123"}},
	}

	cir := ClipResolver{}

	for _, r := range tests {
		_ = clipboard.WriteAll(r.readValues)
		res, err := cir.Resolve(r.inputList)

		if r.errExpected == "" {
			assert.Equal(t, r.expected, res)
		} else {
			assert.NotNil(t, err)
			assert.Empty(t, res)
		}

	}

}

func TestClipInputReader_Read(t *testing.T) {
	tests := []struct {
		inputLIst   *ResolverConfig
		readValues  string
		expected    []string
		errExpected string
	}{
		{
			inputLIst: &ResolverConfig{
				ArgSep: " ",
				Keys:   nil,
			},
			readValues:  "a b c d",
			expected:    []string{"a", "b", "c", "d"},
			errExpected: "",
		},
		{inputLIst: &ResolverConfig{
			Keys:   []string{"-t", "-x"},
			ArgSep: "|",
		}, readValues: "a|b|123", expected: []string{"-t", "a", "-x", "b", "123"}},
	}
	cir := ClipResolver{}

	for _, r := range tests {
		_ = clipboard.WriteAll(r.readValues)
		res, err := cir.Resolve(r.inputLIst)

		if r.errExpected == "" {
			assert.Equal(t, r.expected, res)
		} else {
			assert.NotNil(t, err)
			assert.Empty(t, res)
		}

	}
}
*/
