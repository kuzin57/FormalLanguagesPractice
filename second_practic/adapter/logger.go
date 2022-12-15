package adapter

import (
	"fmt"
	"os"
)

func newLogger(path string) (*logger, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return &logger{
		file: file,
	}, nil
}

func (l *logger) info(configInfo configurationInfo, conf configuration) {
	l.file.WriteString(
		fmt.Sprintf(
			"%d) (%c -> %s, %d) %s\n",
			configInfo.number,
			conf.terminal,
			fmt.Sprintf("%s%c%s", conf.expression[:conf.position], '.', conf.expression[conf.position:]),
			conf.startIndex,
			configInfo.method,
		),
	)
}

func (l *logger) printD(d int) {
	l.file.WriteString(fmt.Sprintf("D%d:\n", d))
}

func (l *logger) printEmptyLine() {
	l.file.WriteString("\n")
}
