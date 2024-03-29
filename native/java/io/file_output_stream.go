package io

import (
	"github.com/murakmii/gojiai/vm"
)

func init() {
	class := "java/io/FileOutputStream"

	vm.NativeMethods.Register(class, "initIDs", "()V", vm.NopNativeMethod)

	vm.NativeMethods.Register(class, "writeBytes", "([BIIZ)V", func(thread *vm.Thread, args []interface{}) error {
		fos := args[0].(*vm.Instance)
		bytes := vm.JavaByteArrayToGo(args[1].(*vm.Instance), int(args[2].(int32)), int(args[3].(int32)))
		fd := fos.GetField("fd", "Ljava/io/FileDescriptor;").(*vm.Instance)

		_, err := fd.AsFile().Write(bytes)
		return err
	})
}
