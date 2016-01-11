## Übung: Key-Value Store
Schreibe einen kleines Programm `kv`, mit dem Du Schlüssel-Werte
Paare in einer Datei speichern und abfragen kannst.

Installieren:
```shell
cd /path/to/go-workshow
go install session01/kv
``` 

Setzen von Werten:
```shell
kv name=Mancke vorname=Sebastian alter=42
```

Abfragen bestimmter Werte:
```shell
kv name vorname
> name=Mancke
> vorname=Sebastian
```

Abfragen aller Werte:
```shell
kv
> name=Mancke
> vorname=Sebastian
> alter=42
```