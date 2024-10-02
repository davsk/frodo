#!/bin/zsh

# "init.zsh"
# by David Lynn Skinner
# for Davsk Ltd Co

# Init script assumes nothing.

# case OS
case "$OSTYPE" in
    darwin*)
        echo Update Darwin MacOS
        sudo softwareupdate -i -a -R
        brew install mas
        ;;
    linux*)
        echo Update Linux
        ;;
    dragonfly*|freebsd*|netbsd*|openbsd*)
        echo Update
        ;;
    *)
        echo $OSTYPE "is Not Supported!"
        exit
        ;;
esac

if [ ! -x $(whence -p brew) ]; then
    echo "brew not found! Please install brew."
    exit 1
else
    #brew update and upgrade packages
    brew update
    brew upgrade

    # insta;; sass
    if [ ! -x $(whence -p sass) ]; then
        brew install sass/sass/sass
    fi

    # update submodules
    git submodule update --remote --merge

    # update golang
    sudo ./update-golang/update-golang.sh

    echo "Update Hugo"
    CGO_ENABLED=1 go install -tags extended github.com/gohugoio/hugo@latest
    #cd ../WealthInnovationsCA/WealthInnovationsCA-Hugo
    #go get -u
    #hugo mod get -u
    #nlm update
fi
