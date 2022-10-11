@echo off

set OUTPUT_NAME=gs_editor
set BUILD_DIR=%~dp0build\x86\
set TOOLS_DIR=%~dp0tools\

@echo build tools...
@cd %TOOLS_DIR%
for /f "tokens=*" %%i in ('dir /a:d /b') do (
	if not exist %TOOLS_DIR%%%i\%%i.exe (
		@echo   - %TOOLS_DIR%%%i
		@cd %TOOLS_DIR%%%i
		@go build
	)
)

@echo cleaning build directory...
if exist %BUILD_DIR% (
	@rmdir /s /q %BUILD_DIR% 2> NUL
)

@echo build app...
@cd %~dp0
set GOARCH=386
@go build -buildmode=exe -ldflags="-H windowsgui" -o %BUILD_DIR%%OUTPUT_NAME%.exe

@echo copy assets...
@cd %~dp0
@xcopy /y /q /c /d %~dp0assets\x86\*.* %BUILD_DIR% 2> NUL
@xcopy /y /q /c /d %~dp0assets\*.* %BUILD_DIR% 2> NUL

:close
@pause
@exit
