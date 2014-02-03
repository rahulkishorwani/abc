"newcluster"
=======

"newcluster" is a repository which has go files for making the interconnections between number of servers(individual server as one cluster) connected to each other & has test file for testing the validity of operations performed by this package.

Language Used
=====
     Go language of Google
Hierarchy
=====
    ```newcluster/src/mycluster/
                              mycluster.go
                              mycluster_test.go
                              serverlist0.xml
                              serverlist1.xml
                              serverlist2.xml
                              serverlist3.xml
                              serverlist4.xml
     ```     

Libraries Used
=====
System:- strconv,strings,fmt,sync,os,os/exec,syscall,time,math/rand,encoding/xml.

Other:-myluster,zmq4.

Requirement
=====
   Requirement zmq4 repository.
   Install it by using the command
   ```go get github.com/pebbe/zmq4```
   

Install
=====
```go get github.com/rahulkishorwani/newcluster```

Run
=====
In the command prompt in mycluster folder run following command
```go test```
This will run only one test. To have multiple tests uncomment the other test functions in mycluster_test.go

Constraints
=====
   Program have been constrained by the configuration(servelist) xml files. For current server configurations, it works 
for 5 servers. If we want to change the the number of servers then we have to change each xml file to accomodate those much of server entries & you have to change the corresponding number of servers in mycluster_test.go & must add those much of serverlist xml files.

Use
=====
   This code can be useful to build a cluster interface to communicate with other clusters.

Rights
=====
   This code is open source. So anyone can refer this.
