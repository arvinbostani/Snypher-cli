#!/bin/bash

BINARY="snypher-linux-amd64"
INSTALL_PATH="/usr/local/bin/snypher"
REPO_URL="https://github.com/arvinbostani/Snypher"

show_help() {
    clear
    echo "==================== SNYPHER HELP ===================="
    echo "Snypher — network sniffer CLI"
    echo ""
    echo "Usage:"
    echo "  snypher list        List interfaces"
    echo "  snypher <iface>     Start sniffer"
    echo "  snypher -r <iface>  Run with flags"
    echo ""
    echo "========================================================"
    echo ""
    read -p "Press Enter to return to menu..."
}

install_snypher() {
    clear
    echo "🔥 Installing Snypher..."

    if [ -f "../dist/$BINARY" ]; then
        echo "Using local binary..."
        sudo cp "../dist/$BINARY" "$INSTALL_PATH"
    else
        echo "Downloading latest binary..."
        sudo curl -L -o "$INSTALL_PATH" \
          "$REPO_URL/releases/latest/download/$BINARY"
    fi

    sudo chmod +x "$INSTALL_PATH"
    echo "✅ Installed! Run 'snypher list'"
    read -p "Press Enter..."
}

uninstall_snypher() {
    clear
    echo "❌ Uninstalling Snypher..."
    sudo rm -f "$INSTALL_PATH"
    echo "Done!"
    read -p "Press Enter..."
}

menu() {
    while true; do
        clear
        echo "==============================================="
        echo "        S N Y P H E R   I N S T A L L E R"
        echo "==============================================="
        echo " Author: arvinbostani"
        echo " License: "
        echo " Repo: $REPO_URL"
        echo "==============================================="
        echo ""
        echo " [1] Install Snypher"
        echo " [2] Uninstall Snypher"
        echo " [3] Help"
        echo " [4] Exit"
        echo ""
        read -p "Choose an option: " opt

        case "$opt" in
            1) install_snypher ;;
            2) uninstall_snypher ;;
            3) show_help ;;
            4) exit 0 ;;
            *) echo "Invalid option"; sleep 1 ;;
        esac
    done
}

menu
