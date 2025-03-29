# quick start

get and run davFileBrowser &emsp;&emsp;&emsp;&emsp;


## 1. What's davFileBrowser

     Based on filebrowser, webdav function is added, which can map local network drive.



## 1. Source mode
	
	git clone https://github.com/yydsapi/davFileBrowser  
	cd davFileBrowser  
	go mod tidy  
    go build
    chmod +x ./davFileBrowser
    ./davFileBrowser
    Parameter reference: https://github.com/filebrowser/filebrowser

    webdav Parameter please modify config.toml at root path
    or reference for more : https://github.com/emersion/go-webdav


### Then open the page address on the run output like this: 2025-03-29 16:27:29 INF WebDAV server started addr=http://0.0.0.0:9997
     
    Map network drive address(use ssl):

    http{s}://{ your webdav address}:{ your webdav port}/{your filebrowser username}

    Username: your filebrowser username
    Password: your filebrowser password

    How to map a network drive:
    Right-click on My Computer, select Map Network Drive, enter the address, 
    select the following two options, click Finish, enter the username and password, and click OK

# Source From :

https://github.com/filebrowser/filebrowser;

https://github.com/emersion/go-webdav


## Contact
mail: vnextcanary@gmail.com  
or visit link: https://yydsapi.com/library/listitem?page=library&category=MUcDljbq_go
https://yydsapi.cn/library/listitem?page=library&category=MUcDljbq_go

## tips:
If you have good projects or suggestions, we can help you realize it
## Special thanks:
@fhst, @kdhly, and all function modules used in the project structure and plug-ins; and other function module not listed;

## Give me a star
If you like or plan to use this project,please give me a star, thank you!

## Donation
If this project makes you feel good, You can donate to the following link to better support the development of this project and the team:
![10](/static/img/donation/alipay.jpg)   <br /><br /> <br />

![10](/static/img/donation/weixin.jpg)    <br /><br /> <br />


## Screenshots ：<br /><br />
#### mainpage 
>![11](/static/img/screenshots/mainpage.jpg)  <br /><br />
#### publish 
>![11](/static/img/screenshots/publish.jpg)  <br /><br />
#### show content 1

