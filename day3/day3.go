package main

import (
	"bufio"
	"fmt"
	"os"
)

func treeSledding(progression int, skip bool) (treeCount int) {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		location := line[counter]
		if string(location) == "#" {
			treeCount++
		}
		if counter+progression < len(line) {
			counter += progression
		} else {
			counter = (counter + progression) - len(line)
		}
		if skip {
			scanner.Scan()
		}
	}
	return
}

func main() {
	fmt.Println("Part one: ", treeSledding(3, false))
	fmt.Println("Part two: ", treeSledding(1, false)*treeSledding(3, false)*treeSledding(5, false)*treeSledding(7, false)*treeSledding(1, true))
	fmt.Println(`
	   *        *        *        __o    *       *
*      *       *        *    /_| _     *
   K  *     K    HO HO HO   O'_)/ \  *    *
  <')____  <')____    __*   V   \  ) __  *
   \ ___ )--\ ___ )--( (    (___|__)/ /*     *
 *  |   |    |   |  * \ \____| |___/ /  *
    |*  |    |   |     \____________/       *
	`)
}
