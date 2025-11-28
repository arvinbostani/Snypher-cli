@echo off
SETLOCAL ENABLEDELAYEDEXPANSION
title SNYPHER - Installer

:menu
cls
echo ==================================================
echo        S N Y P H E R   I N S T A L L E R
echo ==================================================
echo.
echo     creator: arvinbostani (Rvnyx)
echo.
echo     Let's sniff some data... HAHAHA!!
echo.
echo  1) Install Snypher
echo  2) Uninstall Snypher
echo  3) Exit
echo.
set /p choice="Select an option: "

if "%choice%"=="1" goto install
if "%choice%"=="2" goto uninstall
if "%choice%"=="3" exit

goto menu

:install
cls
echo ==================================================
echo                 INSTALLING SNYPHER
echo ==================================================

net session >nul 2>&1
if %errorlevel% neq 0 (
    echo ERROR: Run this installer as Administrator!
    pause
    goto menu
)

echo Building binary...
go build -o snypher.exe cmd/snypher/main.go
if %errorlevel% neq 0 (
    echo Build FAILED!
    pause
    goto menu
)

set TARGET_DIR="%ProgramFiles%\Snypher"

echo Installing to %TARGET_DIR% ...
if not exist %TARGET_DIR% mkdir %TARGET_DIR%

copy /Y snypher.exe %TARGET_DIR% >nul

echo Adding to PATH if missing...
echo %PATH% | find /I %TARGET_DIR% >nul
if %errorlevel% neq 0 (
    setx PATH "%PATH%;%TARGET_DIR%"
)

echo.
echo DONE!
echo Type: snypher -NetMon eth0
pause
goto menu

:uninstall
cls
echo Removing installed files...
set TARGET_DIR="%ProgramFiles%\Snypher"
rd /S /Q %TARGET_DIR% >nul 2>&1

echo Uninstalled successfully.
pause
goto menu
