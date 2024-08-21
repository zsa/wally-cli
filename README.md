# Wally cli (Deprecated)

⚠️⚠️⚠️ Wally is no longer maintained, instead please use our new flashing tool [Keymapp](https://www.zsa.io/flash#download)

Flash your [ZSA Keyboard](https://ergodox-ez.com) the EZ way.

## Getting started
Download the application for your favorite platform from the [release page](https://github.com/zsa/wally-cli/releases).

Note for Linux users, follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install) before running the application.

Note for Mac OS users, the CLI requires libusb to be installed: `brew install libusb`

You can also compile and install Wally using go's package manager, make sure you follow the `Installing dev dependencies` section for your platform below:

```
go get -u github.com/zsa/wally-cli
```

Note: Raspberry pi users using the 32bit version of raspbian should run
```
GOOS=linux GOARCH=arm go get -u github.com/zsa/wally-cli
```

## Automating firmware downloads from the CLI

To get your latest binary all you need to do is got to this url: `https://oryx.zsa.io/{layout ID}/latest/binary`

A few things to note:

- You can also replace the `latest` keyword with a revision ID if you want to get a specific revision.
- You can replace the `binary` keyword with `source`, to download the source code of your layout.
- The URL redirects to our CDN, so you will need to add the -L param to curl. If you use wget it should redirect by default.
- If the revision is not compiled, the endpoint will return a 404
- If the layout is private, the endpoint will return a 401

## Installing dev dependencies
Wally is compatible with Windows, Linux, and macOS. Developing using each platform requires some extra setup:

### Windows
1. Install [TDM GCC](http://tdm-gcc.tdragon.net/download)
2. Setup pkg-config - see [http://www.mingw.org/wiki/FAQ](http://www.mingw.org/wiki/FAQ) "How do I get pkg-config installed?"
3. Grab and install the latest version of libusb [from here](http://sourceforge.net/projects/libusb/files/libusb-1.0/)

### Linux
Follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install)

### macOS
Install libusb using `brew`:

```
brew install libusb
```

### build

```
go build
```


## Sending feedback

As you may have noticed, we do not have GitHub Issues enabled for this project. Instead, please submit all feedback via email to contact@zsa.io — you will find us very responsive. Thank you for your help with Wally!
