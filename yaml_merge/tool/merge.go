package main

import (
	"flag"
	"fmt"
	"github.com/xmdas-link/tools/yaml_merge"
	"log"
	"os"
	"strings"
)

var (
	flagSet      = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	outFile      = flagSet.String("out", "merge/application.yaml", "合并后输出的文件地址")
	configFolder = flagSet.String("folder", "config", "配置文件放置的文件夹地址")
	configFiles  = flagSet.String("files", "application", "要合并的配置文件名，不要带文件类型后缀，用,分割开")
)

func main() {

	flagSet.Parse(os.Args[1:])
	m := yaml_merge.New()

	files := []string{}
	fileNames := strings.Split(*configFiles, ",")
	for _, name := range fileNames {
		files = append(files, fmt.Sprintf("%s/%s.yaml", *configFolder, name))
	}

	for _, file := range files {
		if err := m.AddFile(file); err != nil {
			log.Fatal(err.Error())
		}
	}

	if err := m.SaveAs(*outFile); err != nil {
		log.Fatal(err.Error())
	}
}
