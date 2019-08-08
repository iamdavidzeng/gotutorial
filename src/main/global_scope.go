package main

var b = "G"

func main() {
	n_()
	m_()
	n_()
}

func n_() { print(b) }


func m_() {
	b = "0"
	print(b)
}
