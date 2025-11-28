@echo off
:: =============================================
:: Snypher Interactive Installer
:: Author: arvinbostani
:: Repo: https://github.com/arvinbostani/Snypher-cli
:: =============================================

:: ------------------------------
:: Admin check
:: ------------------------------
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo WARNING: This installer requires administrator privileges.
    echo Please right-click and "Run as administrator".
    pause
    exit /b
)

:: ------------------------------
:: Configuration
:: ------------------------------
set BINARY=snypher-windows-amd64.exe
set INSTALL_DIR=C:\Program Files\Snypher
set INSTALL_PATH=%INSTALL_DIR%\snypher.exe
set REPO_URL=https://github.com/arvinbostani/Snypher-cli

:: ------------------------------
:: Menu Loop
:: ------------------------------
:menu
cls
echo =================================================
echo           S N Y P H E R   INSTALLER
echo =================================================
echo Author: arvinbostani
echo License: MIT
echo Repo: %REPO_URL%
echo =================================================
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

echo Invalid option! Try again.
pause
goto menu

:: ------------------------------
:: Install
:: ------------------------------
:install
cls
echo Installing Snypher...

:: Check for wpcap.dll (Npcap/WinPcap)
if not exist "%SystemRoot%\System32\wpcap.dll" (
    echo WARNING: Npcap or WinPcap is required to capture network traffic.
    echo Please download and install from: https://nmap.org/npcap/
    pause
    goto menu
)

set INSTALL_SUCCESS=0

:: Create install folder if not exists
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

:: Try local binary first
if exist ..\dist\%BINARY% (
    echo Using local binary from project...
    copy /Y ..\dist\%BINARY% "%INSTALL_PATH%" >nul 2>&1
    if %errorlevel%==0 set INSTALL_SUCCESS=1
) else (
    echo Downloading binary from GitHub releases...
    powershell -Command ^
        "try { (New-Object Net.WebClient).DownloadFile('%REPO_URL%/releases/latest/download/%BINARY%', '%INSTALL_PATH%') } catch { exit 1 }"
    if %errorlevel%==0 set INSTALL_SUCCESS=1
)

:: Add folder to PATH if not already
reg query "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path | find "%INSTALL_DIR%" >nul
if errorlevel 1 (
    setx /M Path "%INSTALL_DIR%;%PATH%" >nul
    echo Added "%INSTALL_DIR%" to system PATH. Restart your terminal to use 'snypher'.
)

:: Report result
if %INSTALL_SUCCESS%==1 (
    echo.
    echo INSTALLATION COMPLETE! Snypher is ready to rock.
) else (
    echo.
    echo INSTALLATION FAILED! Something went wrong.
)

pause
goto menu

:: ------------------------------
:: Uninstall
:: ------------------------------
:uninstall
cls
echo Removing Snypher...
if exist "%INSTALL_PATH%" (
    del "%INSTALL_PATH%" >nul 2>&1
    if %errorlevel%==0 (
        echo Snypher uninstalled successfully. All traces removed.
    ) else (
        echo Uninstall failed. Try running this installer as administrator.
    )
) else (
    echo Snypher not found at "%INSTALL_PATH%". Nothing to remove.
)
pause
goto menu

:: ------------------------------
:: Help
:: ------------------------------
:help
cls
echo ================= HELP =================
echo Snypher - Network Sniffer CLI
echo.
echo Usage:
echo   snypher list       - List available network interfaces
echo   snypher <iface>    - Capture traffic on interface
echo   snypher -r <iface> - Run Snypher in a new terminal window
echo.
echo Installer allows easy installation/uninstallation on Windows.
echo Ensure Npcap or WinPcap is installed for packet capture.
echo =========================================
pause
goto menu
