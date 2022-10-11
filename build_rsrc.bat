@echo off

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

@echo build *.syso...
@cd %~dp0
@%TOOLS_DIR%go-winres\go-winres.exe make --in %~dp0winres\winres.json

:close
@pause
@exit
