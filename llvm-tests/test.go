package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/llir/llvm/ir"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	log.Println("Generating code...")

	i32 := types.I32
	a := constant.NewInt(5, i32)
	b := constant.NewInt(13, i32)
	zero := constant.NewInt(0, i32)

	m := ir.NewModule()

	putchar := m.NewFunction("putchar", i32, ir.NewParam("c", i32))

	main := m.NewFunction("main", i32)
	entry := main.NewBlock("")

	tmp1 := entry.NewMul(a, b)
	entry.NewCall(putchar, tmp1) // putchar('A')

	entry.NewRet(zero)

	f, err := os.OpenFile("rand.ll", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(m.String()); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("Compiling to assembler...")
	llc := exec.Command("llc", "rand.ll")
	if err := llc.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Compiling to executable...")
	clang := exec.Command("clang", "-o", "rand.exe", "rand.s")
	if err := clang.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Executing")
	rand := exec.Command("./rand.exe")
	rand.Stdout = os.Stdout
	if err := rand.Run(); err != nil {
		log.Fatal(err)
	}
}
