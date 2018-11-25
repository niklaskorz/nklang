package main

import (
	"io/ioutil"
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
	a := constant.NewInt(i32, 5)
	b := constant.NewInt(i32, 13)
	zero := constant.NewInt(i32, 0)

	m := ir.NewModule()

	putchar := m.NewFunc("putchar", i32, ir.NewParam("c", i32))

	main := m.NewFunc("main", i32)
	entry := main.NewBlock("")

	tmp1 := entry.NewMul(a, b)
	entry.NewCall(putchar, tmp1) // putchar('A')

	entry.NewRet(zero)

	data := []byte(m.String())
	if err := ioutil.WriteFile("rand.ll", data, 0644); err != nil {
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
