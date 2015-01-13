:: Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

setlocal

cd %~dp0

set bpgdec=..\internal\bpg-0.9.5-win32\bpgdec.exe
set bpgenc=..\internal\bpg-0.9.5-win32\bpgenc.exe
set bpgview=..\internal\bpg-0.9.5-win32\bpgview.exe

%bpgdec% -o lena512color.png lena512color.bpg

%bpgdec% -o clock.gif clock.bpg
%bpgdec% -o clock.png clock.bpg

