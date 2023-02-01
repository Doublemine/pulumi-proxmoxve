# ProxmoxVE Resource Provider

The ProxmoxVE Resource Provider lets you manage [ProxmoxVE](http://proxmox.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @doublemine/pulumi-proxmoxve
```

or `yarn`:

```bash
yarn add @doublemine/pulumi-proxmoxve
```


### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/Doublemine/pulumi-proxmoxve/sdk/go/...
```


## Configuration

The following configuration points are available for the `proxmoxve` provider:

- `proxmoxve:pm_api_url` - pm_api_url` - (Required; or use environment variable `PM_API_URL`) This is the target Proxmox API endpoint.
- `proxmoxve:pm_user` - (Optional; or use environment variable `PM_USER`) The user, remember to include the authentication realm
  such as myuser@pam or myuser@pve.
- `proxmoxve:pm_password` - (Optional; sensitive; or use environment variable `PM_PASS`) The password.
- `proxmoxve:pm_api_token_id` - (Optional; or use environment variable `PM_API_TOKEN_ID`) This is
  an [API token](https://pve.proxmox.com/pve-docs/pveum-plain.html) you have previously created for a specific user.
- `proxmoxve:pm_api_token_secret` - (Optional; or use environment variable `PM_API_TOKEN_SECRET`) This uuid is only
  available when the token was initially created.
- `proxmoxve:pm_otp` - (Optional; or use environment variable `PM_OTP`) The 2FA OTP code.
- `proxmoxve:pm_tls_insecure` - (Optional) Disable TLS verification while connecting to the proxmox server.
- `proxmoxve:pm_parallel` - (Optional; defaults to 4) Allowed simultaneous Proxmox processes (e.g. creating resources).
- `proxmoxve:pm_log_enable` - (Optional; defaults to false) Enable debug logging, see the section below for logging details.
- `proxmoxve:pm_log_levels` - (Optional) A map of log sources and levels.
- `proxmoxve:pm_log_file` - (Optional; defaults to "terraform-plugin-proxmox.log") If logging is enabled, the log file the
  provider will write logs to.
- `proxmoxve:pm_timeout` - (Optional; defaults to 300) Timeout value (seconds) for proxmox API calls.
- `proxmoxve:pm_debug` - (Optional; defaults to false) Enable verbose output in proxmox-api-go
- `proxmoxve:pm_proxy_server` - (Optional; defaults to nil) Send provider api call to a proxy server for easy debugging

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/).
