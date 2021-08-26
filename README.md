# Go Mock Workshop

## Why?

The purpose of this repository is to assist someone who is new to testing in go, and specifically how to go about mocking external packages that are consumed by code you wish to test.

Hope you find this useful.

## What?

Both the basic, and beter directory contain code that have the same functionality.
The purpose of these two is to show that when it comes to writing good Go Code, you need to take an aproach where you think about how to test what you write.

Otherwise you can easily end up in a scenario where you need to refactor large amounts of code to be able to unit test your code, and doing that is a massive pain.

## Requirements

- Go installed
- Gomock installed
- Make

## Useful commands

The make file contains relevant commands:
- run basic example: `make run-basic`
- run better example: `make run-better`
- create mocks `make create-mocks`
- test `make test`
