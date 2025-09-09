@echo off
setlocal enabledelayedexpansion

REM Simple build script for Windows
REM Usage: build.bat [target]

set "TARGET=%1"
if "%TARGET%"=="" set "TARGET=all"

echo Building target: %TARGET%

if "%TARGET%"=="all" goto :all
if "%TARGET%"=="build" goto :build
if "%TARGET%"=="clean" goto :clean
if "%TARGET%"=="generate" goto :generate
if "%TARGET%"=="generate-windows" goto :generate-windows
if "%TARGET%"=="generate-windows-ps" goto :generate-windows-ps
if "%TARGET%"=="info" goto :info

echo Unknown target: %TARGET%
echo Available targets: all, build, clean, generate, generate-windows, generate-windows-ps, info
exit /b 1

:all
call :build
goto :eof

:build
echo Building for Windows...
cmake -S . -B build -DCMAKE_BUILD_TYPE=Release
cmake --build build --config Release
echo Build completed. Library: build\libhexlib.dll
goto :eof

:clean
echo Cleaning build directory for Windows...
if exist build rmdir /s /q build
goto :eof

:generate
echo Generating bindings for Windows...
if exist bindings\gen.bat (
    bindings\gen.bat
) else (
    echo gen.bat not found
)
goto :eof

:generate-windows
echo Generating bindings for Windows...
if exist bindings\gen.bat (
    bindings\gen.bat
) else (
    echo gen.bat not found
)
goto :eof

:generate-windows-ps
echo Generating bindings for Windows using PowerShell...
powershell -ExecutionPolicy Bypass -File bindings\gen.ps1
goto :eof

:info
echo Detected OS: Windows
echo Build Directory: build
echo Library Name: libhexlib.dll
echo Library Path: build\libhexlib.dll
echo Remove Command: rmdir /s /q
echo Make Directory Command: mkdir
goto :eof
