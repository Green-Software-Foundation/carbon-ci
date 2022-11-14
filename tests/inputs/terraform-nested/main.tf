terraform {
  required_version = "1.2.5"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.15.1"
    }
  }
  
  backend "azurerm" {}
}

provider "azurerm" {
  features {}
}

locals {
  prefix   = "gsf"
  location = "westus"
  username = "mruser"
  password = "mrpswrd1234!"
  
  other_linux_vms = [
    {
      name     = "${local.prefix}-vm-linux2"
      nic_name = "${local.prefix}-nic-linux2"
      hostname = "${local.prefix}-host-linux2"
    },
	{
      name     = "${local.prefix}-vm-linux3"
      nic_name = "${local.prefix}-nic-linux3"
      hostname = "${local.prefix}-host-linux3"
    }	
  ]
}

resource "azurerm_resource_group" "main" {
  name     = "${local.prefix}-rg"
  location = local.location
}

resource "azurerm_virtual_network" "main" {
  name                = "${local.prefix}-network"
  address_space       = ["10.0.0.0/22"]
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
}

resource "azurerm_subnet" "linux" {
  name                 = "${local.prefix}-linux-subnet"
  resource_group_name  = azurerm_resource_group.main.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = ["10.0.2.0/27"]
}

resource "azurerm_subnet" "windows" {
  name                 = "${local.prefix}-windows-subnet"
  resource_group_name  = azurerm_resource_group.main.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = ["10.0.2.32/27"]
}

resource "azurerm_network_interface" "linux1" {
  name                = "${local.prefix}-nic-linux1"
  resource_group_name = azurerm_resource_group.main.name
  location            = azurerm_resource_group.main.location

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.linux.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_linux_virtual_machine" "main" {
  name                            = "${local.prefix}-vm-linux1"
  resource_group_name             = azurerm_resource_group.main.name
  location                        = azurerm_resource_group.main.location
  size                            = "Standard_D2s_v3"
  admin_username                  = local.username
  admin_password                  = local.password
  disable_password_authentication = false
  network_interface_ids = [
    azurerm_network_interface.linux1.id,
  ]

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }
}

module "core_module" {
  source            = "./modules/core_module"
  rg_name           = azurerm_resource_group.main.name
  location          = azurerm_resource_group.main.location
  admin_username    = local.username
  admin_password    = local.password
  windows_subnet_id = azurerm_subnet.windows.id
  windows_nic_name  = "${local.prefix}-nic-windows1"
  windows_vm_name   = "${local.prefix}-vm-windows1"
  linux_subnet_id   = azurerm_subnet.linux.id
  other_linux_vms   = local.other_linux_vms
}
