# Data generation

Instructions to download and format data sets used by the library

## Country data

```sh
wget https://raw.githubusercontent.com/konstantinstadler/country_converter/master/country_converter/country_data.tsv -O-  | sed "s/\t/,/g" > country_data.csv
```

## Cities

```sh
wget http://download.geonames.org/export/dump/cities500.zip 
unzip cities500.zip
sed -i '1 i\geonameid\tname\tasciiname\talternatenames\tlatitude\tlongitude\tfeature_class\tfeature_code\tcountry_code\tcc2\tadmin1_code\tadmin2_code\tadmin3_code\tadmin4_code\tpopulation\televation\tdem\ttimezone\tmodification_date' cities500.txt
mv cities500.txt cities500.tsv
rm cities500.zip
```

## GeoLite2
GeoLite2 database is distributed by MaxMind, Inc., further instructions can be found here:
https://dev.maxmind.com/geoip/geolite2-free-geolocation-data?lang=en
