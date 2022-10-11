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

	charmapTable = map[int]string{}
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
	var bytes []byte
	str = strings.ReplaceAll(str, "\r\n", "\n")

	for i := 0; i < len(str); {
		seek := 0

		for j := 0; j < 32; j++ {
			seek++

			var foundKeys []int
			if i+j >= len(str) {
				return nil, false
			}

			char := str[i : i+j+1]
			for key, value := range charmapTable {
				if char == value {
					foundKeys = append(foundKeys, key)
					continue
				}
			}

			// 등록된 문자 존재
			if len(foundKeys) == 1 {
				n := foundKeys[0]
				if n > 0xff {
					bytes = append(bytes, byte(((n >> 8) & 0xff)))
					bytes = append(bytes, byte(n&0xff))
				} else {
					bytes = append(bytes, byte(n))
				}
				break
			}

			// 숫자 처리
			if strings.HasPrefix(char, "[") && strings.HasSuffix(char, "]") {
				numberStr := char[1 : len(char)-1]
				if i, err := strconv.ParseInt(numberStr, 16, 8); err == nil {
					bytes = append(bytes, byte(i))
					break
				} else {
					return nil, false
				}
			}
		}

		i += seek
	}

	return bytes, true
}
