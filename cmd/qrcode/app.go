package main

import (
	"github.com/yeqown/go-qrcode/writer/standard"

	"github.com/urfave/cli/v2"
)

var copyright = `Copyright (c) 2018 yeqown

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
`

func newApp() *cli.App { _ = "STUB: not implemented"; return nil }

func generate(ctx *generateContext) error { _ = "STUB: not implemented"; return nil }

// construct a writer based generateContext

type writerMode uint8

const (
	_ writerMode = iota
	writerMode_FILE
	writerMode_TERMINAL
)

// generateContext generate qrcode from context
type generateContext struct {
	text string
	mode writerMode
	FOO  *fileOutputOptions
	TOO  *terminalOutputOptions
}

type fileOutputOptions struct {
	output        string
	outputSuffix  string
	blockSize     uint8
	borders       [4]int
	isCircleShape bool
	transparent   bool
	halftoneImage string
}

func (foo fileOutputOptions) applyOptions() []standard.ImageOption {
	_ = "STUB: not implemented"
	return nil
}

type terminalOutputOptions struct{}

func parseGenerateContextFrom(c *cli.Context) *generateContext {
	_ = "STUB: not implemented"
	return nil
}

// writer mode

// parse borders

func prepareFlags() []cli.Flag { _ = "STUB: not implemented"; return nil }
