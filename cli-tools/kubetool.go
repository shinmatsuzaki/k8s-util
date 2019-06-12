// kubectl(1)のwrapperコマンド、引数を考えるのが面倒だったので作成

// 以下のblodを参考に作成
// https://qiita.com/uchiko/items/5e5cda98ecb510671e56

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/mitchellh/cli"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// kubectl get pod
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Pod struct{}

func (f *Pod) Help() string {
	return "kubetool pod"
}

func (f *Pod) Run(args []string) int {
	out, _ := exec.Command("kubectl", "get", "--all-namespaces", "-owide", "--show-labels", "pods").Output()
	fmt.Printf("%s", out)
	//log.Println("Foo!")
	return 0
}

func (f *Pod) Synopsis() string {
	return "Run kubectl get --all-namespaces -o wide --show-labels pods"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// kubectl get all
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type All struct{}

func (f *All) Help() string {
	return "kubetool all"
}

func (f *All) Run(args []string) int {
	out, _ := exec.Command("kubectl", "get", "--all-namespaces", "-owide", "all").Output()
	fmt.Printf("%s", out)
	//log.Println("Foo!")
	return 0
}

func (f *All) Synopsis() string {
	return "all kubectl get --all-namespaces -o wide all"
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// main
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	// コマンドの名前とバージョンを指定
	c := cli.NewCLI("kubetool", "1.0.0")

	// サブコマンドの引数を指定
	c.Args = os.Args[1:]

	// サブコマンド文字列 と コマンド実装の対応付け
	c.Commands = map[string]cli.CommandFactory{
		"pod": func() (cli.Command, error) {
			return &Pod{}, nil
		},
		"all": func() (cli.Command, error) {
			return &All{}, nil
		},
	}

	// コマンド実行
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
