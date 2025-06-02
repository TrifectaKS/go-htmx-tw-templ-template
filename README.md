install: sudo apt get golang-go
install: go get -t github.com/a-h/templ/cmd/templ@latest~
install: go get github.com/a-h/templ@latest

run: go run main.go

generate template: templ generate