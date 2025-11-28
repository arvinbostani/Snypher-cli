@echo off
title Snypher Installer

set BINARY=snypher-windows-amd64.exe
set INSTALL_PATH=%SystemRoot%\System32\snypher.exe
set REPO_URL=https://github.com/arvinbostani/Snypher

:menu
cls
echo ===============================================
echo        S N Y P H E R   I N S T A L L E R
echo ===============================================
echo  Author: arvinbostani
echo  License:
echo  Repo: %REPO_URL%
echo ===============================================
echo.
echo   [1] Install Snypher
echo   [2] Uninstall Snypher
echo   [3] Help
echo   [4] Exit
echo.
set /p choice="Choose an option: "

if "%choice%"=="1" goto install
if "%choice%"=="2" goto uninstall
if "%choice%"=="3" goto help
if "%choice%"=="4" exit

echo Invalid option!
pause
goto menu

:install
cls
echo Installing Snypher...

if exist ..\dist\%BINARY% (
    echo Using local binary...
    copy /Y ..\dist\%BINARY% "%INSTALL_PATH%" >nul
) else (
    echo Downloading binary...
    powershell -Command "(New-Object Net.WebClient).DownloadFile('%REPO_URL%/releases/latest/download/%BINARY%', '%INSTALL_PATH%')"
)

echo.
echo Installed successfully!
pause
goto menu

:uninstall
cls
echo Removing Snypher...
del "%INSTALL_PATH%" >nul 2>&1
echo Uninstalled!
pause
goto menu

:help
cls
echo ================= HELP =================
echo Snypher - Network Sniffer CLI
echo.
echo Usage:
echo   snypher list
echo   snypher ^<interface^>
echo   snypher -r ^<interface^>
echo =========================================
pause
goto menu
