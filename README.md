# Ansible Inventory To Consul Service Entries

Convert Ansible static inventory

```
[hostgroup1]
1.1.1.1 consul_tag=abc

[hostgroup2]
2.2.2.2
```

into yaml format understood by consul ansible role:

```
consul_services:
- name: hostgroup1
  address: "1.1.1.1"
  tags:
  - abc
- name: hostgroup2
  address: "2.2.2.2"
```

## Example

```
$ go run main.go ./example.ini    
```

```yaml
consul_services:
- address: 10.255.1.11
  name: mysql-slave
- address: 10.255.1.10
  name: mysql-slave
- address: 10.255.1.5
  name: docker-registry
- address: 10.255.1.2
  name: order
  tags:
  - internal
```

## Credits

https://github.com/outten45/aini
