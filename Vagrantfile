Vagrant.configure("2") do |config|
  config.vm.box_check_update = false

  config.vm.define "master" do |master|
    master.vm.box = "ubuntu/xenial64"
    master.vm.hostname = 'salt-master'
    master.vm.network "private_network", ip: "192.168.88.101"
    master.vm.provision "shell", :path => "salt-master.sh"
  end
 
  config.vm.define "minion" do |minion|
    minion.vm.box = "ubuntu/xenial64"
    minion.vm.hostname = 'salt-minion'
    minion.vm.network "private_network", ip: "192.168.88.102"
    minion.vm.provision "shell", :path => "salt-minion.sh"
  end
end
