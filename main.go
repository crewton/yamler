package main

import (
  "log"
  "flag"
  "os"
  "io/ioutil"
  "github.com/smallfish/simpleyaml"
)

func main() {
  var filename = flag.String("f", "", "the filename to parse")
  flag.Parse()

  var data []byte
  var err error
  if *filename == "" {
    log.Fatalln("You need to specify a filename.")
  } else {
    data, err = ioutil.ReadFile(*filename)
    if err != nil {
      log.Fatalln("Error reading file:", err)
    }
  }

  success, err := parseYaml(data)
  if err != nil {
    log.Fatalln("Error parsing yaml:", err)
  }

  if success {
    os.Exit(0)
  } else {
    os.Exit(1)
  }
}

func parseYaml(input []byte) (bool, error) {
  _, err := simpleyaml.NewYaml(input)
  if err != nil {
    return false, err
  }

  return true, nil
}
