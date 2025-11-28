#!/bin/bash

clear

while true; do
    clear
    echo "=================================================="
    echo "        S N Y P H E R   I N S T A L L E R"
    echo "=================================================="
    echo
    echo "      Let's sniff some data HAHAHA!! "
    echo "      Creator : arvinbostani (Rvnyx)"
    echo
    echo " 1) Install Snypher"
    echo " 2) Uninstall Snypher"
    echo " 3) Exit"
    echo
    read -p "Select an option: " choice

    case "$choice" in
        1)
            clear
            echo "Building..."
            if [ "$EUID" -ne 0 ]; then
                echo "Run with sudo!"
                read -p "Press enter..."
                continue
            fi

            go build -o snypher cmd/snypher/main.go

            echo "Installing to /usr/local/bin..."
            mv snypher /usr/local/bin/
            chmod +x /usr/local/bin/snypher

            echo "Done!"
            read -p "Press Enter..."
            ;;

        2)
            clear
            echo "Removing /usr/local/bin/snypher..."
            rm -f /usr/local/bin/snypher
            echo "Done!"
            read -p "Press Enter..."
            ;;

        3)
            exit 0
            ;;

        *)
            echo "Invalid option..."
            sleep 1
            ;;
    esac
done
