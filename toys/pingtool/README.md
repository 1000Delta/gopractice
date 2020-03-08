# pingtool

> The tool use to ping a group of host through linux ping tool.

## hosts file

The file of hosts every valid line should have to fields:

```plain
[hostname] [host]
```

like:

```plain
lo localhost
local 127.0.0.1
```

- `hostname` you should input some word of name.
  - too long hostname will break the format of output.
- `host` you should input the domain or IPv4 address you want to test.
