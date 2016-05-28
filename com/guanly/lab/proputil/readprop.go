package proputil

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func GetProps(proppath string) map[string]string {
	confFile, err := os.Open(proppath)
	if err != nil {
		log.Fatal("properties file not find.")
		panic(err)
	}
	defer confFile.Close()

	rd := bufio.NewReader(confFile)
	m := map[string]string{}

	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		kv := strings.Split(line, "=")
		if len(kv) != 2 {
			continue
		}
		m[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	return m
}
