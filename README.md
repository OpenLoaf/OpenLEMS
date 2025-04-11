

# EMS

## 命名规范
### Golang类型缩写
Enums: 
>E,e

Interface: 
>I,i

Struct: 
>S,s

Function: 
>F,f


### 文件命名
包名：{项目第一个字母小写}_包名 
>PS: base/device -> device需要改名为 b_device

文件名：{大类型}_{小类型(可选)...}_{Golang类型缩写(小写)}_{Golang类型缩写(小写)(可选)...}.go  
>PS: type_e.go type_pcs_i.go type_pcs_exec_f.go


### 代码命名
枚举:
>{Golang类型缩写(大写)}{Name}

接口:
>{Golang类型缩写(大写)}{Name}

结构体:
> {Golang类型缩写(大写)}{Name}


## 运行
### 参数

--driver-path=resources/driver --web=true

## 编译
### 创建镜像
docker build -t ems-go .

### 启动容器
docker run -v .:/root/work -ti ems-go /bin/bash


## 环境部署
### influxdb1
https://docs.influxdata.com/influxdb/v1/introduction/install/

```bash
sudo apt install influxdb-client
```

### influxdb2
https://docs.influxdata.com/influxdb/v2/install/#install-influxdb-as-a-service-with-systemd