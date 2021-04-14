package main

import (
	"fmt"
	"io"
)

var (
	// Version 版本
	Version string
	// CommitHash git 提交 hash 值
	CommitHash string
	// BuildTime 构建时间
	BuildTime string
)

// ShowVersion 输出版本信息
func ShowVersion(w io.Writer) {
	fmt.Fprintf(w,
		`Version
--------------------------------------------------------------
Version:        %s
Commit Hash:    %s
Build Time:     %s
--------------------------------------------------------------
`,
		Version, CommitHash, BuildTime,
	)
}
