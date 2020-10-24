package tail

import(
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Command(script string) {
	path, _ := os.Getwd()
	fmt.Println(path)

	bootCmd := fmt.Sprintf("%s/%s", path, script)

	fmt.Println("bootCmd: ", bootCmd)

	denoCmd := exec.Command("bash", "-c", bootCmd)

	stdout, _ := denoCmd.StdoutPipe()
	denoCmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	denoCmd.Wait()
}
