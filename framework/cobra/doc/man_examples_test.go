package doc_test

import (
	"bytes"
	"fmt"

	"github.com/jamesluo111/core_web/framework/cobra"
	"github.com/spf13/cobra/doc"
)

func ExampleGenManTree() {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "my test program",
	}
	header := &doc.GenManHeader{
		Title:   "MINE",
		Section: "3",
	}
	cobra.CheckErr(doc.GenManTree(cmd, header, "/tmp"))
}

func ExampleGenMan() {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "my test program",
	}
	header := &doc.GenManHeader{
		Title:   "MINE",
		Section: "3",
	}
	out := new(bytes.Buffer)
	cobra.CheckErr(doc.GenMan(cmd, header, out))
	fmt.Print(out.String())
}
