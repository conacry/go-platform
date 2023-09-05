package generator

var (
	lettersForRandom = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandomDefaultStr() string {
	return RandomStr(15)
}

func RandomStr(length int64) string {
	randomLetters := make([]rune, length)
	for i := range randomLetters {
		randomLetterIndex := RandomNumber(0, int64(len(lettersForRandom)-1))
		randomLetters[i] = lettersForRandom[randomLetterIndex]
	}
	return string(randomLetters)
}
