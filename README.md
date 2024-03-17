# Golang Intern Assignment 2

## Install migration 
```
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
```

```
    sudo apt-get update
```
```
    sudo apt-get install migrate
```


## Migrate database
Pull docker postgres:
```
    make run-postgres
```

Run migration:
```
    make migrate
```

## Run application
```
    go run .
```

## Test API 
Using Thunderclient extension in Visual Studio Code: https://www.thunderclient.com/ <br>

First, fetch data from https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_week.geojson by run POST 
localhost:3000/fetchgeo/{number of record you want to fetch} <br>
Data will be saving to postgres database <br>

GET all earthquake record: 
```
localhost:3000/earthquake/all
```

Apply multiple filter by place to earthquake api by run GET: localhost:3000/earthquake?limit=<your_limit>&offset=<your_offset>&place=<your_place> <br>

Example: <br>
* Apply pagination and filtered by place:
  ```
  localhost:3000/earthquake?limit=20&offset=1&place=7 km SW of Holtville, CA
  ```
* Test with default query:
  ```
  localhost:3000/earthquake
  ```
* Filtered by id: localhost:3000/earthquake/{id}, example:
```
  localhost:3000/earthquake/96
```
Get request by hours ago to now by run GET: <br>
localhost:3000/getrequestbyweek/{hours}<br>
Example: 
```
localhost:3000/getrequestbyhour/3
```
Get request by weeks ago to now by run GET: <br>
localhost:3000/getrequestbyweek/{week}<br>
Example: 
```
localhost:3000/getrequestbyweek/2
```




