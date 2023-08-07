Sure! Below is a README format with instructions on how to use Ansible to build and install your Go program and its corresponding systemd service:

**README.md**

# Building and Installing Your Go Program with Ansible

This README provides instructions on how to use Ansible to build your Go program and install it as a systemd service on a target server.

## Prerequisites

1. Ensure that Ansible is installed on your local machine.

2. The target server should have Go installed and properly configured.

## Step 1: Prepare Your Go Program

1. Clone or copy your Go program to a local directory on your machine.

2. Ensure that your Go program has a proper `main.go` file and other required Go files.

## Step 2: Set Up the Ansible Playbook

1. Create a new file named `playbook.yml` in the root directory of your Go program.

2. Paste the following content into `playbook.yml`:

```yaml
---
- name: Build and Install Go Program
  hosts: your_target_server
  remote_user: your_remote_user
  become: yes

  vars:
    go_program_name: "your_program_name"    # Replace with your Go program name
    go_program_path: "/path/to/your/program"  # Replace with the absolute path to your Go program

  tasks:
    - name: Install Go (if not already installed)
      raw: test -e /usr/local/go || (curl -s https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz | tar -C /usr/local -xzf -)

    - name: Set Go environment variables
      environment:
        PATH: "/usr/local/go/bin:{{ ansible_env.PATH }}"
        GOPATH: "/opt/go"

    - name: Build Go Program
      command: "go build -o {{ go_program_name }} {{ go_program_path }}/main.go"
      args:
        chdir: "{{ go_program_path }}"
      changed_when: false

    - name: Create systemd service
      template:
        src: systemd/your_program.service.j2   # Replace with the path to your systemd service template
        dest: /etc/systemd/system/{{ go_program_name }}.service
        owner: root
        group: root
        mode: '0644'
      notify:
        - Reload systemd

  handlers:
    - name: Reload systemd
      systemd:
        name: "{{ go_program_name }}"
        state: reloaded
```

3. In the `vars` section, replace `your_program_name` with your actual Go program name, and set the `go_program_path` to the absolute path of your Go program on the target server.

## Step 3: Create a Systemd Service Template

1. Create a new directory named `templates` in the root directory of your Go program.

2. Inside the `templates` directory, create a new file named `your_program.service.j2` (replace `your_program` with your Go program name).

3. Paste the following content into `your_program.service.j2`:

```ini
[Unit]
Description=Your Go Program
After=network.target

[Service]
User={{ remote_user }}
Group={{ remote_user }}
WorkingDirectory={{ go_program_path }}
ExecStart={{ go_program_path }}/{{ go_program_name }}
Restart=always

[Install]
WantedBy=multi-user.target
```

4. Save the file.

## Step 4: Run the Ansible Playbook

1. Ensure that your target server is reachable and accessible from your local machine.

2. Run the Ansible playbook using the following command:

```bash
ansible-playbook -i "your_target_server_ip_or_hostname," playbook.yml
```

Replace `your_target_server_ip_or_hostname` with the IP address or hostname of your target server. The comma after the server name is important to indicate that it is a single target.

3. Ansible will execute the tasks defined in the playbook, including installing Go (if not already installed), building your Go program, and creating the systemd service.

4. After the playbook completes successfully, your Go program should be built and installed as a systemd service on the target server.

## Step 5: Manage Your Go Program Using Systemd

You can use systemd commands on the target server to manage your Go program:

- Start your program: `sudo systemctl start your_program_name`
- Stop your program: `sudo systemctl stop your_program_name`
- Restart your program: `sudo systemctl restart your_program_name`
- Check the status of your program: `sudo systemctl status your_program_name`

That's it! Your Go program is now built and installed on your target server, managed by systemd. You can easily update or manage your program using Ansible and systemd commands.
