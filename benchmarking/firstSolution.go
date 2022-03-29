package benchmarking

import "os"

func LengthOfFile(f string, bufferSize int) (int, error) {
	f, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	count := 0
	for {
		buffer := make([]byte, bufferSize)
		num, err := f.Read(buffer)
		count += num
		if err != nil {
			break
		}
	}
	return count, nil
}
