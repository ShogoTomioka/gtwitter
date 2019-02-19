package lib

func FormatCreatedAt(createdAt string) string {
	var time []byte
	time = []byte(createdAt)
	return string(time[:20])
}
