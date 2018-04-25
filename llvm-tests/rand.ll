declare i32 @putchar(i32 %c)

define i32 @main() {
; <label>:0
	%1 = mul i32 5, 13
	%2 = call i32 @putchar(i32 %1)
	ret i32 0
}
