# democapt01

## Installation

### Installation from source

1. Download source
   ```bash
   $ git clone https://github.com/wanglei-ok/democapt01.git
   ```
1. CD to Directory
   ```bash
    $ cd democapt01/
   ```
1. Test Run
   ```bash
    $ GO111MODULE=on go run *.go
   ```

## service

1. 创建服务配置文件

    在远程服务器上编写服务配置，放在 /etc/systemd/system/ 中
    ```bash
    [Unit]
    Description=democapt01 service
    After=network.target
    
    [Service]
    User=wallace
    WorkingDirectory=/home/wallace/src/democapt01/
    # Execute `systemctl daemon-reload` after ExecStart= is changed.
    Environment="LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/home/wallace/src/democapt01/hksdk/linux/lib/"
    ExecStart=/home/wallace/src/democapt01/democapt01
    [Install]
    WantedBy=multi-user.target
    ```
    修改User、WorkingDirectory、Environment、ExecStart中的路径信息,确保正确

1. 启动服务
    ```bash
    # 每一次修改ExecStart都需执行
    systemctl daemon-reload
    # 启动
    systemctl start democapt01.service
    # 查看状态
    systemctl status democapt01.service
    ```
   
1. Open Browser
    ```bash
    open web browser http://127.0.0.1:10086/public work chrome, safari, firefox


## API documentation

See the [API docs](/doc/api.md)