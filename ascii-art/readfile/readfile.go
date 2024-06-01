package readfile

import "os"

//read the contents of the file
func ReadFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)
	
	if err != nil {
		return "Error Reading File", err
	}
	return string(fileContent), nil
}
