# **Infra Patch Manager**

_A **CLI tool** - which helps to maintain the remote infra patching. There is a scenario, a distributed system is running on multiple servers and an upgrade/change required for all these server's application(s) so there is an repeated activity will be performed to logging in each and every server's application to apply an upgrade._

_Here, **infrapm** (Infra Patch Manager) does the work. it helps to maintain an upgrade at centralized place and apply at required destination with help of few sub-commands._

### **Installation | Setup**

---

_**infrapm** has master/minion(agent) concept to work over servers. One infrapm agent will be running on each server(s) as a service with specific port [default: 8008], infrapm master will stand where all application servers are reachable._

_To setup the infrapm, taking reference server (Linux : x86_64). Download the latest infrapm release version with specific os-architure package. After extracting the package two executable files mainly as_

- `infrapm`
- `infrapm_agent`

_and LICENSE, README.md, conf/remotes.json(for reference) etc._

#### SETUP AGENT
***

_Copy the `infrapm_agent` file to destination application server. Same agent setup steps will repeated for mutliple application servers. Copy on linux as_

```
scp infrapm_agent user@HOSTIP:/home/user/
```

_After copying create a working directory for agent as_

```
mkdir -p /opt/infrapm
mv infrapm_agent /opt/infrapm/
```

_Create a linux service file as_

```
vi /opt/infrapm/infrapm.service
```

_Use service file reference as: alter the running port if needed_

```
[Unit]
Description=Infra Patch Manager

[Service]
Type=simple
WorkingDirectory=/opt/infrapm
ExecStart=/opt/infrapm/infrapm_agent --port 8008
ExecStop=

[Install]
WantedBy=multi-user.target
```

_To enable the service for linux system. refer these steps as_

```
cd /etc/systemd/system
ln -s /opt/infrapm/infrapm.service
systemctl daemon-reload
```

_Start the infrapm agent as_

```
systemctl start infrapm.service
systemctl status infrapm.service
```

_Make sure the infrapm agent working directory lies as follows_

    infrapm_agent
    infrapm.service
    resources/
        assets/
        patch/
        rollback/


#### SETUP MASTER
***

_Copy the `infrapm` file at destination where application server(s) are reachable. And create a file as **conf/remotes.json** there as below, repeat servers and server applications configuration as needed._

```
[
  {
    "agent_address": "x.x.x.x:8008",
    "type": "xyz",
    "name": "remote_server_1",
    "apps": [
      {
        "type": "abc",
        "name": "application_1",
        "source": "/opt/application_1/path",
        "service": "application_1",
        "port": "8000"
      },
      {
        "type": "xyz",
        "name": "application_2",
        "source": "/opt/application_2/path",
        "service": "application_2",
        "port": "8002"
      },
    ]
  }
]
```

_To test the infrapm master functionality by running as_

> ./infrapm --help

_Expected output will be_

```
Infra-Patch-Manager contains the following subcommands set.

        remote          | list or search a remote detail with reachablity
        rights          | read/write rights check on a remote's application(s)
        upload          | upload a patch to remote
        extract         | untaring a tar.gz file on relative remote
        apply           | applying a patch to relative remote application(s)
        verify          | helps to validate an applied patch
        exec            | helps to execute commands on remote(s)

```

_Make sure the infrapm master working directory lies as follows_

    conf/
        remotes.json
    infrapm

### **infrapm | Explore**

---

_As completed the setup of infrapm application in two phase agent (repeated installed if remote servers are more) and master (once) installation. To use the application, always redirect to working directory of infrapm master. Explore the infrapm funcationality by subcommands help. here are few subcommands example as_

> ./infrapm --help

> ./infrapm remote -all

> ./infrapm remote

> ./infrapm remote -name xyz -ping

> ./infrapm rights

> ./infrapm rights --remote-all -app-all

> ./infrapm extract -remote-all -list

_Completed **:)**_
