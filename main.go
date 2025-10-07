// Command yaml2json allows to convert YAML on STDIN to JSON on STDOUT.
//
// Install
//
//	go install github.com/dolmen/yaml2json
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	goyaml "gopkg.in/yaml.v2"
)

func main() {
	err := translate(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	os.Exit(0)
}

func translate(in io.Reader, out io.Writer) error {
	input, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	var data interface{}
	err = goyaml.Unmarshal(input, &data)
	if err != nil {
		return err
	}
	input = nil
	fixMaps(&data)

	output, err := json.Marshal(data)
	if err != nil {
		return err
	}
	data = nil
	_, err = out.Write(output)
	return err
}

func fixMaps(pIn *interface{}) {
	switch in := (*pIn).(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(in))
		for k, v := range in {
			ks, isString := k.(string)
			if !isString {
				ks = fmt.Sprint(k)
			}
			fixMaps(&v)
			m[ks] = v
		}
		*pIn = m
	case []interface{}:
		for i := len(in) - 1; i >= 0; i-- {
			fixMaps(&in[i])
		}
	}
}
