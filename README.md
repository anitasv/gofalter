gofalter
========

Lisp in Go
https://www.ee.ryerson.ca/~elf/pub/misc/micromanualLISP.pdf

But like the title suggests it is not the whole truth :) so don't use this lisp, except from the perspective of implementing what is
on this paper!

export GOPATH=$HOME/go

mkdir -p $GOPATH/src/github.com/anitasv

cd $GOPATH/src/github.com/anitasv

git clone github.com/anitasv/gofalter

go install github.com/anitasv/gofalter

$GOPATH/bin/gofalter

