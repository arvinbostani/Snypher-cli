## 👁️ Snypher CLI: Terminal Network Sniffer

Snypher CLI is a powerful, terminal-based network traffic analysis tool built with Go. It allows you to monitor network activity on a specified interface in real-time, displaying packet details and flagging potentially **suspicious traffic**.

---

## ✨ Features

* **Real-Time Monitoring:** Display live network traffic directly in your terminal.
* **Interface Selection:** Easily specify the network adapter to monitor (e.g., `eth0`, `Wi-Fi`).
* **Packet Details:** Shows source/destination IP, port, protocol, and payload size.
* **Suspicious Activity Flagging:** Alerts the user to packets matching pre-defined heuristics for unusual or potentially malicious traffic.
* **Cross-Platform:** Supports Linux and Windows systems.

---

## 💻 Installation

Snypher includes self-contained scripts to automatically build and install the executable to a location accessible from your system's `PATH`.

### Prerequisites

You must have **Go** installed on your system to run the installation scripts.

### Linux / macOS

1.  Navigate to the `install` directory.
    ```bash
    cd install
    ```
2.  Run the installation script using `sudo`, as root permissions are required to move the executable to a system directory and for raw socket access during sniffing.
    ```bash
    sudo sh install.sh
    ```

### Windows

1.  Navigate to the `install` directory.
    ```bash
    cd install
    ```
2.  Run the installation script.
    ```bash
    install.bat
    ```

---

## 🚀 Usage

Once installed, you can run the sniffer from any directory. You **must** specify the network interface using the `-i` flag.

### Linux / macOS

You need **root privileges (`sudo`)** to capture packets on most Linux systems.

| Interface Type | Command |
| :--- | :--- |
| **Wired** (Example) | `sudo snypher -r <interface>` |
| **list of interface** (Example) | `sudo snypher list` |

### Windows

| Interface Type                   | Command                  |
|:---------------------------------|:-------------------------|
| **Wired** (Example)              | `snypher -r <interface>` |
| **list of interfaces** (Example) | `snypher list`           |

> ℹ️ **Note:** Interface names on Windows are case-sensitive and must match the names listed in your Network Connections.

---

## 🚨 Suspicious Traffic Detection

Snypher flags traffic as **SUS** based on a set of internal heuristics designed to identify activity commonly associated with network scanning, unauthorized access attempts, or unusual protocol behavior.

The current detection logic includes, but is not limited to, the following indicators:

* **High Port Counts:** Packets targeting an **unusually high number of diverse destination ports** in a short period (suggestive of a port scan).
* **Malformed Packets:** Packets that violate standard protocol specifications (e.g., TCP SYN packets with payload data).
* **Unusual Internal Communication:** Unexpected traffic patterns or protocols seen on reserved private IP ranges.
* **Excessive Failed Attempts:** A large volume of failed connection attempts to a single destination.

When a packet is flagged as suspicious, it will be highlighted in the terminal output with a clear **[SUS]** indicator.

---

## 🏗️ Project Structure

The project follows standard Go practices for structure and modularity:

| Directory/File | Description |
| :--- | :--- |
| `cmd` | Main package containing the application's entry point and CLI setup. |
| `sniff` | The core package containing the logic for packet capture, parsing, and suspicious traffic analysis. |
| `ui` | Package dedicated to rendering the Terminal User Interface (TUI) for real-time display. |
| `install` | Contains `install.sh` (Linux/macOS) and `install.bat` (Windows) for automated setup. |
| `go.mod` | Go module definition file for dependencies. |