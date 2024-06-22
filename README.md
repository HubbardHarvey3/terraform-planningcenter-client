# terraform-planningcenter-client
This repository contains a Go client library tailored for interacting with the Planning Center API.

The client was originally built to allow the [Planning Center Terraform Provider](https://github.com/HubbardHarvey3/terraform-provider-planningcenter) to use the Planning Center API.

## Using the Client

To import the client, you can run
`go get github.com/HubbardHarvey3/terraform-planningcenter-client`

If you only require certain packages you can import those specific packages with

`import (github.com/HubbardHarvey3/terraform-planningcenter-client/people)`

### Authentication

Currently, the client uses 2 environment variables to authenticate your account with Planning Center.

To setup authentication, you will need your Planning Center [Personal Access Tokens](https://api.planningcenteronline.com/oauth/applications)

After creating your "Personal Access Token," you can export them as environment variables

On Linux

`export PC_APP_ID=<Your Application ID> `

`export PC_SECRET_TOKEN=<Your Secret>`

On Windows (Powershell)

`$env:PC_APP_ID=<Your Application ID> `

`$env:PC_SECRET_TOKEN=<Your Secret>`

## License

The repo is licensed under Mozilla Public License Version 2.0.  For more information, you can read the `LICENSE` file in the root of the repo.

## Contributing

Contributions are welcome.  However, this is a hobby project and I am unable to dedicate a lot of time to its maintenance.  
If you have a problem or an idea for an enhancement, please create an issue and I will tackle it as soon as possible.


## TODO
- [ ] Need to setup Mock API Responses so testing isn't dependent on a specific account
- [x] Implement Phone Numbers 
