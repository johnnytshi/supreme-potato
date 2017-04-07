package pod

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/supreme-potato/db"
)

func rebaseRepo(r string) {
	fmt.Println("# Rebasing Spec Repo")
	cmd := exec.Command("bash", "-c", "git pull --rebase")
	cmd.Dir = r
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("# Worked")
}

func Run(r string) {
	fmt.Println("### Proccessing Podspec at", r)

	rebaseRepo(r)
}

func SetupLocal() {
	db.DB().CreateTable("Podspecs")
}
