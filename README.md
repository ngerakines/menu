# menu

Menu is minimalistic file format and command line utility to describe and interacting with pairings of applications and cookbooks.

# File Format

```json
{
  "time":1385056233,
  "artifacts":[
    {
      "id":"menu-webapp",
      "location":"file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.war",
      "version":"1.0.0"
    }
  ],
  "cookbooks":[
    {
      "location":"https://github.com/ngerakines/menu-webapp-cookbook"
    }
  ],
  "deploy": [
    {
      "id": "menu-webapp-tomcat",
      "artifact": "menu-webapp",
      "type": "tomcat"
    }
  ]
}
```

# Usage

To create a new menu item:

    $ menu create --artifact-id=menu-webapp --artifact-version=1.0.0 --artifact-location=file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.war --cookbook=https://github.com/ngerakines/menu-webapp-cookbook

To describe a menu item:

    $ menu show file:///var/releases/1385056233.menu
    time: 1385056233
    artifacts:
        menu-webapp 1.0.0 file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip
    cookbooks:
        https://github.com/ngerakines/menu-webapp-cookbook

To list only artifacts or cookbooks for a specific menu item:

    $ menu artifacts file:///var/releases/1385056233.menu
    menu-webapp 1.0.0 file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip

    $ menu cookbooks file:///var/releases/1385056233.menu
    https://github.com/ngerakines/menu-webapp-cookbook

To list all menu items for an artifact:

    $ menu search --artifact-id=menu-webapp file:///var/releases/
    file:///var/releases/1385056233.menu

To create a local-deploy script for a menu item:

    $ menu local-deploy file:///var/releases/1385056233.menu
    #!/bin/sh
    MENU_BASE=`pwd`
    # Clone and start cookbook 'https://github.com/ngerakines/menu-webapp-cookbook' master
    git clone https://github.com/ngerakines/menu-webapp-cookbook e52223dd99ab5dadbab64e38871596b5027841ef
    cd e52223dd99ab5dadbab64e38871596b5027841ef
    vagrant up
    # Get artifact 'file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.war'
    cd $MENU_BASE
    cp /var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.war 35a19d45d918a4866a55437e323fa310f1d33650
    # Deploy 'menu-webapp-tomcat' (tomcat)
    curl --upload-file 35a19d45d918a4866a55437e323fa310f1d33650 "http://admin:password@localhost/manager/deploy?path=/&update=true"


# License

Copyright (c) 2013,2014 Nick Gerakines <nick@gerakines.net>, Chris Antenesse <chris@antenesse.net>

This project and its contents are open source under the MIT license.
