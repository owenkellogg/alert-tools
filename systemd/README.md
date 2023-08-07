Sure! Here's a `systemd/README.md` file with instructions for installing and running your Go program under systemd:

**systemd/README.md**

```
# Running Your Go Program with Systemd

Systemd is a powerful system and service manager for Linux systems. It can be used to manage services, including your Go program, by creating a systemd service unit.

## Prerequisites

1. Your Go program should be compiled and ready to run. Ensure that it's working correctly when executed manually.

2. Your Go program should have the necessary permissions to access the required resources (e.g., database, Rocket.Chat webhook, etc.).

## Step 1: Create a systemd Service Unit File

1. Create a new file with the `.service` extension in the `/etc/systemd/system/` directory. For example, you can name it `your-program.service`:

```bash
sudo nano /etc/systemd/system/your-program.service
```

2. Paste the following content into the file, replacing the placeholders with the appropriate values:

```
[Unit]
Description=Your Go Program
After=network.target

[Service]
User=your_username
Group=your_group
WorkingDirectory=/path/to/your/program
ExecStart=/path/to/your/program/your-program

[Install]
WantedBy=multi-user.target
```

- `Description`: Provide a description for your service.
- `User`: Replace with the username you want the service to run as.
- `Group`: Replace with the group name for the user running the service.
- `WorkingDirectory`: Replace with the absolute path to your program's directory.
- `ExecStart`: Replace with the absolute path to your compiled Go program.

## Step 2: Reload Systemd and Enable the Service

1. Reload the systemd manager configuration to load the new service unit:

```bash
sudo systemctl daemon-reload
```

2. Enable the service to start on boot:

```bash
sudo systemctl enable your-program
```

## Step 3: Start and Manage the Service

1. Start your program using systemd:

```bash
sudo systemctl start your-program
```

2. Check the status of your program and view its logs:

```bash
sudo systemctl status your-program
journalctl -u your-program
```

3. Stop and restart your program using systemd:

```bash
sudo systemctl stop your-program
sudo systemctl restart your-program
```

## Uninstalling

1. To stop and disable the service:

```bash
sudo systemctl stop your-program
sudo systemctl disable your-program
```

2. To remove the service unit file:

```bash
sudo rm /etc/systemd/system/your-program.service
```

3. Finally, reload the systemd manager configuration:

```bash
sudo systemctl daemon-reload
```

That's it! Your Go program is now managed by systemd and will start automatically on system boot. You can use systemd's powerful features for service management, such as restart policies, logging, and monitoring.
```

Make sure to replace the placeholders with the actual values specific to your Go program, such as the path to your compiled binary, your username, group, etc. Also, update the service description to match your program.
