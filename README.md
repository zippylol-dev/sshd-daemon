SSHD Daemon for ADB-based shell access

This daemon sets up a minimal SSH server accessible as the shell user on your device, allowing interactive sessions directly into the device.


---

🧭 Setup Instructions

1. Complete the setup for the Custom-daemon-for-adb:

Make sure that daemon is working before you proceed.


2. Deploy the SSH daemon files:

Clone this repository into /sdcard.

Copy the contents of prebuilt(or manually compile the client and server using go) into /data/local/tmp:

cp -r /sdcard/SSHD-daemon/prebuilt/* /data/local/tmp


                                                     ---
                                                     ⚡ Start the SSHD:                                                                                        Use the Daemon_call function defined in your Termux ~/.bashrc:                                                                                                 Daemon_call /data/local/tmp/wrapper
                                                     If you see failed to bind to port, you can ignore this — it still works!                                                                                                                                            ---                                                  
🔑 Connect to the SSH server:                                                                             Copy the ssh_client binary into your Termux home directory:

./myssh

Optional WAN/LAN access: Before running the client, you can set a remote host:

export MYSSH_SERVER="[ip]:port" # the ip has to be surrounded in [] but port doesnt have to be
./myssh


---

🧠 Notes

Requires the ADB daemon to be running.

The SSHD stays alive until the device is restarted.

Uses localhost by default.



---

📜 License

MIT License


---
