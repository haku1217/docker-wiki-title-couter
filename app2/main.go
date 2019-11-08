
package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
	"compress/gzip"
	"io"
)
func LinesInFile(fileName io.Reader) []string {
   // Create new Scanner.
   scanner := bufio.NewScanner(fileName)
   result := []string{}
   // Use Scan.
   for scanner.Scan() {
       line := scanner.Text()
       // Append line to result.
       result = append(result, line)
   }
   return result
}
func contain(s []string, e string) int {
	var count int
    for _, v := range s {
		if regexp.MustCompile(e).Match([]byte(v)) {
			count++
		}
	}
	return count
}
// func check_regexp(reg, str string) {
//     fmt.Println(regexsp.MustCompile(reg).Match([]byte(str)))
// }

func main() {
	args := strings.Split(os.Args[1], ",")
	// Loop over lines in file.
	file, _ := os.Open("jawiki-latest-all-titles.gz")
	defer file.Close()

	gzipReader, _ := gzip.NewReader(file)
    defer gzipReader.Close()

	lines := LinesInFile(gzipReader)
	for _, s := range args {
		fmt.Println(s,":",contain(lines, s))
	}
   // Get count of lines.
}
