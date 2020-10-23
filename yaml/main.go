package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type test struct {
	Name    string    `yaml:"name"`
	Sarapan []sarapan `yaml:"sarapan"`
	Sekolah sekolah   `yaml:"sekolah"`
}

type sarapan struct {
	Makan string `yaml:"makan"`
	Minum string `yaml:"minum"`
}

type sekolah struct {
	Key map[string]sekolahName `yaml:",inline"`
}

type sekolahName struct {
	Name string `yaml:"name"`
}

func main() {
	t := test{}
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(t)

	for key, val := range t.Sekolah.Key {
		fmt.Println(key)
		fmt.Println(val)
	}
	fmt.Println(t.Sekolah.Key)
}
