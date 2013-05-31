package gedis

import (
	"bytes"
	"strconv"
)

/*
 Redis protocoal

 Redis uses a very simple text protocol, which is binary safe.

 *<num args> CR LF
 $<num bytes arg1> CR LF
 <arg data> CR LF
 ...
 $<num bytes argn> CR LF
 <arg data>
*/

// Writes a string as a sequence of bytes to be send to a Redis
// instance, using the Redis Bulk format.
func WriteBulk(bulk string) []byte {
	bulk_len := strconv.Itoa(len(bulk))

	// '$' + len(string(len(bulk))) + "\r\n" + len(bulk) + "\r\n"
	n := 1 + len(bulk_len) + 2 + len(bulk) + 2

	bytes := make([]byte, n)

	bytes[0] = '$'

	j := 1

	for _, c := range bulk_len {
		bytes[j] = byte(c)
		j++
	}

	bytes[j] = '\r'
	bytes[j+1] = '\n'
	j += 2

	for _, c := range bulk {
		bytes[j] = byte(c)
		j++
	}

	bytes[j] = '\r'
	bytes[j+1] = '\n'

	return bytes
}

// Writes a sequence of strings as a sequence of bytes to be send to a
// Redis instance, using the Redis Multi-Bulk format.
func WriteMultiBulk(cmd string, args ...string) []byte {
	var buffer bytes.Buffer

	buffer.WriteByte('*')
	buffer.WriteString(strconv.Itoa(1 + len(args)))
	buffer.WriteString("\r\n")

	buffer.Write(WriteBulk(cmd))

	for _, elem := range args {
		buffer.Write(WriteBulk(elem))
	}

	return buffer.Bytes()
}