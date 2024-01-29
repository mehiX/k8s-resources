/*
Read YAML input (ex: the output of `helm template`) and output the resources needed
based on the resources declarations and replicas
*/
package main

import (
	"github.com/mehix/k8s-resources/cmd"
)

func main() {
	cmd.Execute()
}
