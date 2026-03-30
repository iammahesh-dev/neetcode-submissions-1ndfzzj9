type linkedliststack struct {
	stack []byte
}
func Constructor() linkedliststack {
	return linkedliststack{
		stack: make([]byte, 0),
	}
}
func (this *linkedliststack) push(val byte){
	this.stack = append(this.stack, val)
}
func (this *linkedliststack) pop()byte {
	val:= this.stack[len(this.stack) - 1]
	this.stack= this.stack[:len(this.stack) - 1]
	return val
}

func (this *linkedliststack) empty() bool {
	return len(this.stack)==0
}

func isValid(s string) bool {
    stack := Constructor()
	closeOpenMap := map[byte]byte{')': '(', ']': '[', '}':'{'}
	for i:=0; i<len(s); i++ {
		if open, ok := closeOpenMap[s[i]]; ok {
			if stack.empty(){
				return false
			}
			val:= stack.pop()
			if ok && open != val{
				return false
			}
		} else {
			stack.push(s[i])
		}
	}
	return stack.empty()
}
