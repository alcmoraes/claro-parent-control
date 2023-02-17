# YIP
### Youth Internet Punisher
---

A tool that allows you to punish young fuc*ers by managing your home router configuration.

## Requirements

As of now, this tool works only with routers from the **Claro** ISP. A model called `HGB10R-02`

You can check your model by accessing the router's web interface, and looking at the inspector's console.

![Screenshot](https://github.com/alcmoraes/yip/blob/master/static/router-inspector.png)

## Configuration
Put the `.yip.json` inside your `$HOME` folder and configure the credentials used to access your router.
You can check your browser inspector to retrieve your username and password used to authenticate.

## Usage


### List devices connected to the DHCP
```shell
$ yip listDevices
```

Will output all devices connected to the DHCP, with their mac address, and their hostname.

```bash
===== DEVICES =====
64:64:4A:39:E3:38 - MiWiFi-R4CM
24:A1:60:3F:51:37 - ESP_3F5137
F8:54:B8:94:C5:24 - XBOX <- aka. The victim
64:A2:00:01:E8:6B - RedmiNote8-RedmiNote
88:66:5A:1E:AA:38 - BR0C02FM02JMD6M
===================
```

### Adds a mac address to the blacklist (aka. the naughty list)
```shell
$ yip filterDevice 64:A2:00:01:E8:6B 24:A1:60:3F:51:37 ...
```

### Removes a mac address from the blacklist
```shell
$ yip unfilterDevice 64:A2:00:01:E8:6B 24:A1:60:3F:51:37 ...
```