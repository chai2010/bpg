:: Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

setlocal

cd %~dp0

del /Q                 lena512color.png
..\tools\bpgdec.exe -o lena512color.png lena512color.bpg

