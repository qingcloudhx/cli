package template

var (
	Main = `
package main
import (
        "flag"
        "fmt"
        "github.com/qingcloudhx/core/support/log"
        "io/ioutil"
        "os"

        _ "github.com/qingcloudhx/core/data/expression/script"
        "github.com/qingcloudhx/core/engine"
)

var (
        fileJson      = flag.String("conf", "", "app config")
        cfgJson       string
        cfgCompressed bool
)

func main() {

        flag.Parse()
        if *fileJson != "" {
                data, err := ioutil.ReadFile(*fileJson)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Failed to read config file: %s\n", *fileJson)
                        os.Exit(1)
                } else {
                        cfgJson = string(data)
                }
        }
        cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
        if err != nil {
                fmt.Fprintf(os.Stderr, "Failed to create engine: %v\n", err)
                os.Exit(1)
        }
        log.Init(cfg.Log)
        e, err := engine.New(cfg)
        if err != nil {
                fmt.Fprintf(os.Stderr, "Failed to create engine: %v\n", err)
                os.Exit(1)
        }

        code := engine.RunEngine(e)


        os.Exit(code)
}
`
)
