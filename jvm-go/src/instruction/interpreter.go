package instruction

import (
	"fmt"
	"github.com/classfile"
	"github.com/runtimedata"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := runtimedata.NewThread()

	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *runtimedata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *runtimedata.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := NewInstruction(opcode)
		inst.FetchOperands(&reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(*frame)
	}
}