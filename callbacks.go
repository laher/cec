package cec

// #include <libcec/cecc.h>
import "C"

import (
	"log"
	"unsafe"
)

//export logMessageCallback
func logMessageCallback(c unsafe.Pointer, msg *C.cec_log_message) C.int {
	log.Println(C.GoString(msg.message))

	conn := (*Connection)(c)
	conn.messageReceived(C.GoString(msg.message))
	return 0
}

//export keyPressed
func keyPressed(c unsafe.Pointer, code *C.cec_keypress) C.int {
	log.Println("code", code)

	conn := (*Connection)(c)
	conn.keyPressed(int(C.int(code.keycode)))
	return 0
}

//export commandReceived
func commandReceived(c unsafe.Pointer, msg *C.cec_command) C.int {
	// log.Printf("%v", msg)

	conn := (*Connection)(c)
	cmd := &Command{
		initiator:   uint32(msg.initiator),
		destination: uint32(msg.destination),
		ack:         int8(msg.ack),
		eom:         int8(msg.eom),
		opcode:      int(msg.opcode),
		// parameters: todo
		opcode_set:       int8(msg.opcode_set),
		transmit_timeout: int32(msg.transmit_timeout),
		Operation:        opcodes[int(msg.opcode)],
	}
	conn.commandReceived(cmd)

	return 0
}
