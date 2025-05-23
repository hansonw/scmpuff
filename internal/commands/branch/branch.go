package branch

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// CommandBranch generates the command handler for `scmpuff branch`
func CommandBranch() *cobra.Command {
	var branchCmd = &cobra.Command{
		Use:   "branch",
		Short: "List git branches with numbers",
		Run: func(cmd *cobra.Command, args []string) {
			out := gitBranchOutput()
			lines := processBranches(out)
			for i, line := range lines {
				if i > 0 {
					fmt.Println(line)
				} else {
					fmt.Print(line)
					if len(lines) > 1 {
						fmt.Println()
					}
				}
			}
		},
	}
	return branchCmd
}

func gitBranchOutput() []byte {
	out, err := exec.Command("git", "branch", "--no-color").Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func processBranches(out []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(out))
	var head string
	var others []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "*") {
			head = strings.TrimSpace(strings.TrimPrefix(line, "*"))
		} else {
			others = append(others, strings.TrimSpace(line))
		}
	}
	var results []string
	n := 1
	if head != "" {
		if strings.HasPrefix(head, "(") {
			results = append(results, "* "+head)
		} else {
			results = append(results, fmt.Sprintf("* [%d] %s", n, head))
			n++
		}
	}
	for _, b := range others {
		if b == "" {
			continue
		}
		results = append(results, fmt.Sprintf("  [%d] %s", n, b))
		n++
	}
	return results
}
