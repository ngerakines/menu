# menu

Menu is minimalistic file format and command line utility to describe and interacting with pairings of applications and cookbooks.

# File Format

```json
{
  "time":1385056233,
  "artifacts":[
    {
      "id":"www-resources",
      "location":"s3://company-dist/www-resources/www-resources-1.0.0.zip",
      "version":"1.0.0"
    },
    {
      "id":"company-product",
      "location":"s3://company-dist/company-product/company-product-1.0.0.zip",
      "version":"1.0.0"
    }
  ],
  "cookbooks":[
    {
      "location":"github.com/company/webapp-cookbook"
    },
    {
      "location":"github.com/company/webapp-resources-cookbook"
    }
  ]
}
```

# Usage

To create a new menu item:

    $ menu create --artifact-id=www-resources --artifact-location=s3://company-dist/www-resources/www-resources-1.0.0.zip --cookbook=github.com/company/webapp-cookbook

To describe a menu item:

    $ menu show s3://company-releases/1385056233.menu
    time: 1385056233
    artifacts:
        www-resources s3://company-dist/www-resources/www-resources-1.0.0.zip
    cookbooks:
        github.com/company/webapp-cookbook

To list only artifacts or cookbooks for a specific menu item:

    $ menu artifacts s3://company-releases/1385056233.menu
    www-resources 1.0.0 s3://company-releases/1385056233.menu

    $ menu cookbooks s3://company-releases/1385056233.menu
    github.com/company/webapp-cookbook

To list all menu items for an artifact:

    $ menu search --artifact-id=www-resources s3://company-releases/
    s3://company-releases/1385056233.menu
