# menu

Menu is minimalistic file format and command line utility to describe and interacting with pairings of applications and cookbooks.

# File Format

```json
{
  "time":1385056233,
  "artifacts":[
    {
      "id":"menu-webapp",
      "location":"file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip",
      "version":"1.0.0"
    }
  ],
  "cookbooks":[
    {
      "location":"https://github.com/ngerakines/menu-webapp-cookbook"
    }
  ]
}
```

# Usage

To create a new menu item:

    $ menu create --artifact-id=menu-webapp --artifact-version=1.0.0 --artifact-location=file:///var/artifacts/com/socklabs/menu-webapp/menu-webapp-1.0.0.zip --cookbook=https://github.com/ngerakines/menu-webapp-cookbook

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

# License

Copyright (c) 2013 Nick Gerakines <nick@gerakines.net> and Chris Antenesse <chris@antenesse.net>

This project and its contents are open source under the MIT license.
