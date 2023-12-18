enable-git-hooks:
	git config --local include.path ../.gitconfig
	chmod ug+x .githooks/*

# make run
# make run file=docs/examples/hello
run:
	go run src/main.go $(file)

test:
	go test ./...