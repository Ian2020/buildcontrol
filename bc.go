package main

import (
    "errors"
    "fmt"
    "json"
    "os"
    "strings"
)

func get_platform() (platform string, err error) {
    for _, envString := range os.Environ() {
        if strings.Index(envString, "windir") != -1 {
            windir := strings.SplitAfter(envString, "=")[1]
            _, err := os.Stat(windir)
            if err == nil {
                platform = "windows"
            }
            break
        }
        if strings.Index(envString, "???") != -1 {
            platform = "linux"
            break
        }
    }

    if platform == "" {
        platform = "unknown"
        err = errors.New("Unable to determine platform")
    }

    return platform, err
}

func main() {
    //Use env vars to work out if we're Linux/Windows
    //Load our builld definitions
    //Feed in the platform to decide what we're going to call next
    //Each step of the build should deinfed for each platform we support, or declared NotImpl


    //Build def should look lik
/*


step
alias/nickname (shortname/alt name by which it can be invoked)
prerequisites
    [linux::dnf::group'dev tools'::version::DONTCARE]
    [linux::dnf::libreoffice]
    [windows::powershellv4]
    [windows::choco::libreoffice]
thing to call by platform
    [windows::psake::build] or 
    [linux::make::build]

As we know they're using psake/make we should automatically add those requirements

*/
    var platform, err = get_platform()
    if err != nil {
        fmt.Printf("%s\n", err.Error())
        os.Exit(1)
    }
    fmt.Printf("Platform: %s\n", platform)


    //Read a JSON file from .bc directory
    type PlatformSpec struct {
        Platform        string
        Spec            []string
    }

    type BuildControlConfig struct {
        Step            string
        Alias           string
        Description     string
        Prerequisites   []PlatformSpec
        Actions         []PlatformSpec

    }

    //Load some bytes

    var bcc BuildControlConfig
    err := json.Unmarshal(b, & bcc)
}
