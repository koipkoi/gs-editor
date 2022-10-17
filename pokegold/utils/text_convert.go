package utils

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed charmap.txt
	charmap string

	charmapTable        = map[int]string{}
	charmapReverseTable = map[string]int{}
)

func init() {
	lines := strings.Split(strings.ReplaceAll(charmap, "\r\n", "\n"), "\n")
	for _, line := range lines {
		keyValue := strings.Split(line, "=")
		if len(keyValue) == 2 {
			key, err := strconv.ParseInt(keyValue[0], 16, 64)
			value := keyValue[1]

			if err == nil {
				charmapTable[int(key)] = value
				charmapReverseTable[value] = int(key)
			}
		}
	}
}

// 전용 인코딩 바이트를 문자열로 치환
func TextDecode(bytes []byte) string {
	var sb strings.Builder

	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		if b >= 0x1 && b <= 0xb {
			mb := (int(b) << 8) + int(bytes[i+1])
			sb.WriteString(charmapTable[mb])
			i++
		} else {
			if ch, has := charmapTable[int(b)]; has {
				sb.WriteString(ch)
			} else {
				sb.WriteString(fmt.Sprintf("[%2x]", b))
			}
		}
	}

	return sb.String()
}

// 문자열을 전용 인코딩 바이트로 치환
func TextEncode(str string) []byte {
	encoded, _ := TryTextEncode(str)
	return encoded
}

// 문자열을 전용 인코딩 바이트로 치환
//
// 인코딩된 바이트와 성공 유무를 함께 반환함
func TryTextEncode(str string) ([]byte, bool) {
	bytes := make([]byte, 4096)
	if length, ok := TryTextEncodeBuffered(bytes, str); ok {
		return bytes[0:length], true
	}
	return nil, false
}

// 문자열을 전용 인코딩 바이트로 치환하여 버퍼에 기록
func TextEncodeBuffered(buffer []byte, str string) int {
	length, _ := TryTextEncodeBuffered(buffer, str)
	return length
}

// 문자열을 전용 인코딩 바이트로 치환하여 버퍼에 기록
//
// 인코딩된 바이트의 길이와 성공 유무를 함께 반환함
func TryTextEncodeBuffered(buffer []byte, str string) (int, bool) {
	str = strings.ReplaceAll(str, "\r\n", "\n")
	strLength := len(str)
	bufferIndex := 0

	var foundKeys []int
	seek := 0
	for i := 0; i < strLength; {
		seek = 0

		for j := 0; j < 32; j++ {
			seek++

			foundKeys = nil
			if i+j >= strLength {
				return -1, false
			}

			char := str[i : i+j+1]
			if value, ok := charmapReverseTable[char]; ok {
				foundKeys = append(foundKeys, value)
			}

			// 등록된 문자 존재
			if len(foundKeys) == 1 {
				n := foundKeys[0]
				if n > 0xff {
					buffer[bufferIndex] = byte(((n >> 8) & 0xff))
					buffer[bufferIndex+1] = byte(n & 0xff)
					bufferIndex += 2
				} else {
					buffer[bufferIndex] = byte(n)
					bufferIndex++
				}
				break
			}

			// 숫자 처리
			if strings.HasPrefix(char, "[") && strings.HasSuffix(char, "]") {
				numberStr := char[1 : len(char)-1]
				if i, err := strconv.ParseInt(numberStr, 16, 8); err == nil {
					buffer[bufferIndex] = byte(i)
					bufferIndex++
					break
				} else {
					return -1, false
				}
			}
		}

		i += seek
	}

	return bufferIndex, true
}
