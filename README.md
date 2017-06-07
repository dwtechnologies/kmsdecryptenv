# kmsdecryptenv

This program will decrypt ENV variables encrypted with AWS automatically.

It has two modes `auto` and `marker`.

It will decrypt all the envs at once - to maximize performance.

## Auto mode

With auto mode it will try and figure out if your ENV is base64 encoded or not. But since normal vars can look to be encoded we have to assume
that if the decryption fails that it was becasue the string is not an KMS encrypted string. So please mind that this mode can be a bit more risky
since KMS errors will be treated as "This string is not KMS encrypted".

## Marker mode

With marker mode we will decrypt all variables that are called anything with a specified marker.
The default marker is `"KMS_DECRYPT"`, but this can be changed by setting the `"KMS_MARKER"` ENV variable.

It will then decrypt the value of the ENV and set it to a new ENV variable with the same name but with the marker removed.

So for example `TEST_KMS_DECRYPT` would become `TEST`. `VAR_KMS_DECRYPT_1` would become `VAR1` and so forth.

## Authing against AWS

The program will auth against AWS in the following manner

- Using AWS standard ENV variables for auth.
- Using AWS standard credentials file.
- Using the servers EC2 Role (recommended way).

## Possible configuration

The program uses the following ENV vars for configuration

    KMS_AWS_REGION = Region to use for the KMS service
    KMS_MARKER = Marker to use for finding vars to decrypt
    KMS_AUTO = Set to true if auto mode should be enabled (KMS_MARKER is ignored if this is set to true)
    KMS_OUTPUT = Define the output. Supports the following vars:
        {KEY} = The key
        {VAL} = Unencrypted Value
        {LF} = UNIX newline (LF)
        {CRLF} = Windows newline (CRLF)
        {TAB} = Tab

        So to mimic the default output you would set KMS_OUTPUT to "{KEY} {VAL}{LF}".

## Install

Either download any of the binaries provided in the list below or build yourself.
The binaries are self contained and have no dependencies.

You will find them here:
[github.com/dwtechnologies/kmsdecryptenv/releases/tag/1.0.0](https://github.com/dwtechnologies/kmsdecryptenv/releases/tag/1.0.0)

## Bulding

Please clone this to $GOPATH/src/github.com/dwtechnologies folder. Either using `git` or by using `go get -u github.com/dwtechnologies/kmsdecryptenv`.

If you don't have a go env installed, please visit golang.org for info on how to set it up or just grab one of the binary files above.

```bash
cd $GOPATH/src/github.com/dwtechnologies.com/kmsdecryptenv
go get
go build
```

You can also use the supplied Makefile for building and cross-compiling. You will need to have run `go get` beforehand, becasue the Makefile will not do this step.

## Usage

Please see example.sh for a linux/unix working example.

Note that the program doesn't actually set any vars but will output it as `KEY=VALUE` pairs seperated with `\n` (newline).

So using bash you could easily set your ENV by running `export $(./kmsdecryptenv)`

## Performance

The program has an very small memory and CPU footprint. For example running the program outside AWS (from our Stockholm office to eu-west-1) and trying to decode 30 KMS encrypted strings generates the following

```bash
time ./example.sh
        0.44 real         0.16 user         0.03 sys
```