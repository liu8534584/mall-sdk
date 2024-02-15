package utils

import (
	"bytes"
	"strings"
)

func LineToHump(str string) string {
	if !strings.Contains(str, "_") {
		return str
	}
	strList := strings.Split(str, "_")
	var result string
	for k, v := range strList {
		if k == 0 {
			result = v
		} else {
			result += strings.ToUpper(v[:1]) + v[1:]
		}
	}
	return result
}

func StrPad(Input string, PadLength int, PadString string, PadType int) string {

	var leftPad, rightPad = 0, 0

	numPadChars := PadLength - len(Input)

	if numPadChars <= 0 {
		return Input
	}

	var buffer bytes.Buffer

	buffer.WriteString(Input)

	switch PadType {

	case 0:

		leftPad = numPadChars

		rightPad = 0

	case 1:

		leftPad = 0

		rightPad = numPadChars

	case 2:

		rightPad = numPadChars / 2

		leftPad = numPadChars - rightPad

	}

	var leftBuffer bytes.Buffer

	/* 左填充：循环添加字符*/

	for i := 0; i < leftPad; i++ {

		leftBuffer.WriteString(PadString)

		if leftBuffer.Len() > leftPad {

			leftBuffer.Truncate(leftPad)

			break

		}

	}

	/* 右填充：循环添加字符串*/

	for i := 0; i < rightPad; i++ {

		buffer.WriteString(PadString)

		if buffer.Len() > PadLength {

			buffer.Truncate(PadLength)

			break

		}

	}

	leftBuffer.WriteString(buffer.String())

	return leftBuffer.String()

}
