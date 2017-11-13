Open Service Broker Actions (Demo) UI
====

# Overview

This is an example of how a UI can be built dynamically based on a returned swagger from an OSB server.

## Pre-requisites
OSB Actions demo server up and running

## How to use the UI
Open ui/index.html in your browser
Enter in any "Service", "Plan", "Instance" names. 
* E.g "service1", "plan1", "instance1".

Enter in a discovery URL from the OSB server
* E.g "http://localhost:8080/v2/actions/swagger" (for backup & restore)
* "http://localhost:8080/v2/actions/swagger2" (for pause/start/stop/restart)

Create some backups
* click "createBackup" and fill in the details.

View your backups
* click "getBackups" to get a list of backups you've created

Create some restores
* click on "createRestore" and fill in the details

I think you get the rest...