@echo off

@echo clean tools...
set TOOLS_DIR=%~dp0tools\
@cd %TOOLS_DIR%
for /f "tokens=*" %%i in ('dir /a:d /b') do (
	if exist %TOOLS_DIR%%%i\%%i.exe (
		@echo clean %TOOLS_DIR%%%i\%%i.exe
		@cd %TOOLS_DIR%%%i
		@del %%i.exe 2> NUL
	)
)

@echo remove build directory
@rmdir /s /q %~dp0build\ 2> NUL

:close
@pause
@exit
