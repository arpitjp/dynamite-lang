enable-git-hooks:
	git config --local include.path ../.gitconfig
	chmod ug+x .githooks/*

# make run
# make run file=src/1.txt
run:
	go run src/main.go $(file)