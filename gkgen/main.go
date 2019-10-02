/*
Copyright © 2019 NAME HERE AliasYermukanov

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
  "github.com/AliasYermukanov/gkgen/cmd"
)

func main() {
  //viper.AutomaticEnv()
  //gosrc := utils.GetGOPATH() + afero.FilePathSeparator + afero.FilePathSeparator
  //pwd, err := os.Getwd()
  //if err != nil {
  //  logrus.Error(err)
  //  return
  //}
  //if !strings.HasPrefix(pwd, gosrc) {
  //  logrus.Error("The project must be in the $GOPATH folder for the generator to work.")
  //  return
  //}
  cmd.Execute()
}
