# Steps to use SUB Interface Data Source and Resource Provider

## Datasource
SUB details can be retervied based on SUB Name & Device ID, below is example

#### Terraform File
```
data "fmc_devices" "device" {
    name = "10.100.0.218"
}

data "fmc_security_zones" "my_security_zone" {
  name = "TEST2"
}

output "security_zone" {
  value = data.fmc_security_zones.my_security_zone
}
 

data "fmc_device_sub_interface" "my_fmc_device_sub_interface" {
   name = "TenGigabitEthernet0/0"
    device_id = data.fmc_devices.device.id
}

output "my_fmc_device_sub_interface" {
  value = data.fmc_devices.device
}

```
#### After applying you will see below output
![image](https://github.com/ketan-gote/ddd-example/assets/23295769/5597dde5-2b10-4a74-9c5d-dfdd9c706c2e)


## Resource Provider
Using resource provider you can create, update & destroy the Sub Interface

#### Create Sub Interface 

```
data "fmc_devices" "device" {
    name = "10.100.0.218"
}

data "fmc_security_zones" "my_security_zone" {
  name = "TEST2"
}

output "security_zone" {
  value = data.fmc_security_zones.my_security_zone
}
 
resource "fmc_device_sub_interface" "my_fmc_device_sub_interface" {
    device_id = data.fmc_devices.device.id
    name = "TenGigabitEthernet0/0"
    type = "SubInterface"
    description = "DescriptionUpdated"
    sub_interface_id = 3
    if_name = "IFNameUpdated"
    mtu = 1504
    priority = 4
    management_only = false
    security_zone_id = data.fmc_security_zones.my_security_zone.id
    vlan_id = 4
    enabled = true
    ipv4 {
      static {
        address = "3.3.3.3"
        netmask = 4
      }
       dhcp {
        enable_default_route_dhcp = true
        dhcp_route_metric = 0
      }
    }
}


```
#### After applying you will see below output
![image](https://github.com/ketan-gote/ddd-example/assets/23295769/cf935b45-3863-47ff-9b05-353ea93d2af4)

#### Update Sub Interface
Make some change to resource provider and apply, in below example we have update the Description, Logical Name, MTU, Priority & VLAN ID

![image](https://github.com/ketan-gote/ddd-example/assets/23295769/16355183-bb06-4307-9bee-6d680c826823)

#### Destroy Sub Interface
Using destroy command Sub Interface will destroy the Sub Interface which was created

![image](https://github.com/ketan-gote/ddd-example/assets/23295769/6a25a257-5413-4cd2-b40e-8a861f11e30a)






