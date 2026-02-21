# ns-exporter
Nightscout exporter to InfluxDB

It reads data from Nightscout via v3-api. See {{ns-uri}}/api3-docs for details.
Then it pushes this data to an [InfluxDB](https://www.influxdata.com/).

## usage variants:
### 1. inline
```
go build
./ns-exporter
```

### 2. docker

#### Option 1 - local build

Build it locally:
```
docker build -t welfwalter/ns-exporter .
```

Upload to [welf-walter/welfwalter](https://hub.docker.com/repository/docker/welfwalter/welf-walter/general):
```
docker push welfwalter/ns-exporter:latest
```

#### Option 2 - central build
  Pushes to `master` branch lead to a central build by [a Github Action](https://github.com/welf-walter/ns-exporter/actions).

#### Run

Start the container with your parameters:
```
docker run -d  \
  --name ns-exporter \
  --env NS_EXPORTER_NS_URI=https://abc.eu.nightscoutpro.com \
  --env NS_EXPORTER_NS_TOKEN=$NIGHTSCOUTTOKEN \
  --env NS_EXPORTER_INFLUX_URI=https://eu-central-1-1.aws.cloud2.influxdata.com \
  --env NS_EXPORTER_INFLUX_TOKEN=$INFLUXTOKEN \
  --env NS_EXPORTER_INFLUX_ORG=xxyyzz \
  --env NS_EXPORTER_INFLUX_BUCKET=nightscout \
  welfwalter/ns-exporter:latest
```

## Configuration
Configuration is done via parameters:

	mongo-uri       - MongoDb uri to download from
	mongo-db        - MongoDb database name
	ns-uri          - Nightscout server url to download from
	ns-token        - Nigthscout server API Authorization Token
	limit           - number of records to read
	skip            - number of records to skip
	influx-uri      - InfluxDb uri to download from
	influx-token    - InfluxDb access token
	influx-org      - (optional, default = 'ns') InfluxDb organization to use
	influx-bucket   - (optional, default = 'ns') InfluxDb bucket to use
	influx-user-tag - (optional, default = 'unknown') InfluxDb 'user' tag value to be added to every record - to be able to store multiple users data in single bucket


Parameters also can be provided via env with `NS_EXPORTER_` prefix:

	NS_EXPORTER_MONGO_URI=
	NS_EXPORTER_MONGO_DB=
	NS_EXPORTER_NS_URI=
	NS_EXPORTER_NS_TOKEN=
	NS_EXPORTER_LIMIT=
	NS_EXPORTER_SKIP=
	NS_EXPORTER_INFLUX_URI=
	NS_EXPORTER_INFLUX_TOKEN=
	NS_EXPORTER_INFLUX_ORG=
	NS_EXPORTER_INFLUX_BUCKET=
	NS_EXPORTER_INFLUX_USER_TAG=

So you can choose the data source: direct MongoDB or Nightscout REST API. Supplying required set of parameters will trigger related consumer.
You can even supply both and get from both sources :)

For NS API access you need provide security token. For security reason it is better to go to 'Admin tools' and create special token for NS-Exporter only instead of using admin security key. 
Since exporter only requires read access, creating role with two permissions will be enough:
- api:treatments:read
- api:devicestatus:read

## Presentation

I'm using Grafana dashboard for viewing data. To setup grafana with InfluxDB you need to follow InfluxDB's [instructions](https://docs.influxdata.com/influxdb/v2.3/tools/grafana/).
The sample dashboard can be imported from `grafana.json`. It uses both InfluxQL and Flux datasources for different panels. Some can be omitted, some can be reworker based on other InfluxDB datasource query type. 
Anyway they're provided as samples, for educational purpose :)