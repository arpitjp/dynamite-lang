enable-git-hooks:
	git config --local include.path ../.gitconfig
	chmod ug+x .githooks/*