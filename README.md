# quick start

get and run vnextcanary &emsp;&emsp;&emsp;&emsp;[(中文)](/READMECN.md)

## 1. Source mode
	
	git clone https://github.com/kdhly/vnextcanary.git  
	cd vnextcanary  
	go build  
	chmod a+x vnextcanary  
	./vnextcanary  

### components:

	go get github.com/Sirupsen/logrus  
	go get github.com/gin-contrib/sessions  
	go get github.com/gin-gonic/gin  
	go get github.com/go-sql-driver/mysql  
	go get github.com/mattn/go-sqlite3  
	go get github.com/quasoft/memstore  
	etc  


## 2. executable file 

	git clone https://github.com/kdhly/vnextcanary.git   
	cd vnextcanary/dist  
	unzip vnextcanary/dist/dist.rar  

	windows(x64):  
	cd windows  
	copy vnextcanary/sqlite.db and vnextcanary/documents to this folder  
	run vnextcanary.exe as administrator in cmd window 

	linux(x64):  
	cd linux  
	copy vnextcanary/sqlite.db and vnextcanary/documents to this folder  
	chmod a+x vnextcanary  
	./vnextcanary  

	mac(x64):  
	cd mac  
	copy vnextcanary/sqlite.db and vnextcanary/documents to this folder  
	chmod a+x vnextcanary  
	./vnextcanary  

	arm:  
	cd arm  
	copy vnextcanary/sqlite.db and vnextcanary/documents to this folder  
	chmod a+x vnextcanary  
	./vnextcanary  

### then openpage: http://127.0.0.1 

# configuration file: bbs_config_main_i18n_xxx.json
you can change the IPAddr to 127.0.0.1 to visit it only on your local machine, The default database is sqlite3 , you can import .sql in vnextcanary/dist/mysql directory to your mysql database and change DbType to "mysql" if you want to use mysql.

# Structure and Features:
golang + gin + vue + element ui + i18n;    
plug-ins: ckedit + html5player + blueimp Gallery (music and photo manager)  
database: mysql or sqlite;  
1. Localization of all resources;  
2. Prevent script injections attacks, text filtering,Prevent intrusion attack,XSS attacks;  
3. Open and compact portable design;  
4. This system function was more comprehensive and useful,There are relatively complete sharing functions and attachments when publishing,multimedia capabilities,Can be used as knowledge base and personal electronic notepad and personal media center and simple bbs;  
5. You can set up a private blog;  
6. Menu and blog categories online modification;  
7. User information online modification;  
8. Article and reply online modification;  
9. Basic support for all kinds of mobile browsers  

#### Notice:
1. Users level LV5 (score > 2999) and above can publish pictures, attachments, media, and default management user name:limon,password:password;
2. Multiple administrators can be specified, and those with user allow > 100 are administrators;
3. You can adjust some parameters yourself in bbs_config_main_i18n.json file,For example, to find the password, you must configure the correct email server address, SMTP User name and password, etc;
4. BbsUploadPath must have read and write permission;
5. The image directory you want to display should be in BbsUploadPath+"/Picture/photos", the directory name beginning with "my" is not in the list, but it can still be displayed in the URL
6. Known compatible versions go 1.10 or above,gin v1.4 or above,database: mysql 5.7 or above or sqlite 3;

#### Question:
1. There are still a few bugs and non modular content,But it doesnot affect the experience;
2. Old browsers are not supported because Vue uses the ES6 Promise object feature that older browsers cannot emulate, you can seach some plugin to compatible with it;

## Contact
mail: vnextcanary@gmail.com  
or visit link: http://www.vnextcanary.com/?page=bbs&category=XWdNEvaL_go

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
>![11](/static/img/screenshots/show1.jpg)  <br /><br />
#### show content 2
>![11](/static/img/screenshots/show2.jpg)  <br /><br />
#### display photo 
>![11](/static/img/screenshots/photoshow1.jpg)  <br /><br />
#### manage1 
>![11](/static/img/screenshots/manage1.jpg)  <br /><br />
#### manage2 
>![11](/static/img/screenshots/manage2.jpg)  <br /><br />
