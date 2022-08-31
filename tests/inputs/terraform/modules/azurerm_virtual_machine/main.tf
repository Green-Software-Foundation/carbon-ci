resource "azurerm_network_interface" "nic" {
  name                = var.nic_name
  resource_group_name = var.rg_name
  location            = var.location

  ip_configuration {
    name                          = "internal"
    subnet_id                     = var.subnet_id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_virtual_machine" "vm" {
  name                  = var.vm_name
  location              = var.location
  resource_group_name   = var.rg_name
  vm_size               = "Standard_D2s_v3"
  network_interface_ids = [azurerm_network_interface.nic.id]

  storage_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-focal"
    sku       = "20_04-lts-gen2"
    version   = "latest"
  }

  storage_os_disk {
    name          = "${var.hostname}osdisk"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = var.hostname
    admin_username = var.admin_username
    admin_password = var.admin_password
  }
  
  os_profile_linux_config {
    disable_password_authentication = false
  }
}