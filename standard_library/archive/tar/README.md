## Overview
Package tar implements access to tar archives.

Tape archives (tar) are a file format for storing a sequence of files that can be read and written in a streaming manner. This package aims to cover most variations of the format, including those produced by GNU and BSD tar tools.

## 我的理解
tar提供了对buffer的读写操作，可以利用io.Copy写入到标准输出中