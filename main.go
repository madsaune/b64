package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	decodeIt bool
	str      string
	isUrl    bool
)

func main() {
	flag.BoolVar(&decodeIt, "d", false, "decode string")
	// TODO: Implement base64.URLEncoding.EncodeToString()
	flag.BoolVar(&isUrl, "u", false, "use url encoding")
	flag.StringVar(&str, "s", "", "the string to encode/decode")
	flag.Parse()

	if str != "" {
		if decodeIt {
			ds, err := b64decode(str)
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to decode string", err)
				os.Exit(1)
			}

			fmt.Println(string(ds))
			return
		}

		fmt.Println(b64encode([]byte(str)))
		return
	}

	input, err := read(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	if decodeIt {
		ds, err := b64decode(strings.Join(input, "\n"))
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to decode string", err)
			os.Exit(1)
		}

		fmt.Println(string(ds))
		return
	}

	s := b64encode([]byte(strings.Join(input, "\n")))
	fmt.Println(s)
}

func read(r io.Reader) ([]string, error) {
	var txt []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt = append(txt, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return txt, nil
}

func b64encode(val []byte) string {
	return base64.StdEncoding.EncodeToString(val)
}

func b64decode(val string) ([]byte, error) {
	ds, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, err
	}

	return ds, nil
}
