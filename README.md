# Cisco CMX Mock Server 

This is a simple mock server written in Go (Golang) to mimic a few endpoints of Cisco CMX.
This server can be used for importing location hierarchy and floor plans.

## How to run

```shell
docker-compose up -d
```

## Available Endpoints

API Reference: 
https://www.cisco.com/c/en/us/td/docs/wireless/mse/10-6/api/b_cmx_106_api_reference/configuration.html

### Import maps

This endpoint returns all maps.

```
GET /api/config/v1/maps
```

### Import hierarchy

This endpoint returns the ids of all campuses, buildings, floor and zones.

```
GET /api/config/v1/heterarchy/allUserLevels
```

### Import floor plan images

This endpoint returns an image by image name.

```
GET /api/config/v1/maps/imagesource/:imageName
```


## License

This project is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
