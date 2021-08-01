# DropaFetch - Get information about your Huawei router!
DropaFetch is a project made in Go to get system information about your Huawei router. This tool _may_ work for others brands.

Currently it supports showing the total system memory and usage, uptime and kernel version. More features will be added as soon as possible.

# Downloads
Check [the releases page](https://github.com/princessmortix/dropafetch/releases).

# How to use this?

First, SSH or TelNet in your router, (If you don't know what is the login details of your router, try asking on [Huawei Forum](https://forum.huawei.com/enterprise/en/index.html).), then type `su` and then `shell`. First, make sure you have _actually_ shell access by thing `help`. You should this screen: 

![image](https://user-images.githubusercontent.com/47502554/127074116-0c5b3e28-ebf2-4f92-88f8-3db5ab912601.png)

If you get this output, you're unluckly and don't have access to the shell.
![image](https://user-images.githubusercontent.com/47502554/127074431-2756421d-5d8e-462e-90bb-131b0dc21d66.png)

So, if you got access to it, go to a read-write directory (e.g.: /var) by typing `cd /var`, then create a folder with `mkdir dropafetch`. After that, go to that folder with `cd dropafetch`. Then, try downloading the file with `wget <link>`. If the download fails you'll need to make a FTP server on your PC/Phone (If you don't know how to make one, check [this](https://google.com/search?q=how+to+make+a+ftp+server). After that, type `ftpget <your local ip> -P <the ftp server port> dropafetch-linux-arm` (Make sure that dropafetch is on server root), then type `chmod +x dropafetch-linux-arm` and `./dropafetch-linux-arm`, and finally, the program should run and display your Huawei router system information. 
