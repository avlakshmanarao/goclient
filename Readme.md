## Getting Started

Run Terry's Tacky Car Parts in local machine 
Download/clone gocliet repo which essentially contains following files

 - docker-compose.yml
 - main.go
 - main_test.go
 
 Edit docker-compose.yml file and adjust evironment variables
 - volumes and working_dir should be pointed to the directory where sourcecode downloaded
 - APP_BASE_URL should be Terry's Tacky Car Parts application url running in your local machine. 

 Run docker-compose up command to execute all tests. You can see "TestRegisterPage" test case will get faild and remaining 3 testcases will get passed